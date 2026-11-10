import java.net.InetAddress
import java.net.ServerSocket
import java.net.Socket
import java.nio.file.Files
import java.sql.DriverManager
import org.json.JSONObject

fun main() {
    val path = Files.createTempFile("interview-kotlin-", ".db")
    try {
        DriverManager.getConnection("jdbc:sqlite:$path").use { db ->
            db.createStatement().use { statement ->
                statement.executeUpdate("CREATE TABLE things (name TEXT)")
                statement.executeUpdate("INSERT INTO things VALUES ('Hello, world')")
            }
        }
        DriverManager.getConnection("jdbc:sqlite:$path").use { db ->
            db.createStatement().use { statement ->
                statement.executeQuery("SELECT name FROM things").use { rows ->
                    check(rows.next() && rows.getString(1) == "Hello, world") { "SQLite persistence failed" }
                }
            }
        }
        val loopback = InetAddress.getLoopbackAddress()
        ServerSocket(0, 1, loopback).use { listener ->
            listener.soTimeout = 5000
            Socket(loopback, listener.localPort).use { client ->
                listener.accept().use { server ->
                    client.soTimeout = 5000
                    server.soTimeout = 5000
                    for ((sender, receiver) in listOf(client to server, server to client)) {
                        val message = JSONObject().put("message", "Hello, world").toString()
                        sender.getOutputStream().write("$message\n".toByteArray(Charsets.UTF_8))
                        sender.getOutputStream().flush()
                        val received = receiver.getInputStream().bufferedReader(Charsets.UTF_8).readLine()
                        check(JSONObject(received).getString("message") == "Hello, world") { "JSON over TCP failed" }
                    }
                }
            }
        }
        println("Kotlin environment ready: SQLite persistence and JSON over TCP verified.")
    } finally {
        Files.deleteIfExists(path)
    }
}
