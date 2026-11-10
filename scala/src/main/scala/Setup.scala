import java.net.{InetAddress, ServerSocket, Socket}
import java.nio.charset.StandardCharsets.UTF_8
import java.nio.file.Files
import java.sql.DriverManager
import scala.util.Using.resource

@main def setup(): Unit =
  val path = Files.createTempFile("interview-scala-", ".db")
  try
    resource(DriverManager.getConnection(s"jdbc:sqlite:$path")) { db =>
      resource(db.createStatement()) { statement =>
        statement.executeUpdate("CREATE TABLE things (name TEXT)")
        statement.executeUpdate("INSERT INTO things VALUES ('Hello, world')")
      }
    }
    resource(DriverManager.getConnection(s"jdbc:sqlite:$path")) { db =>
      resource(db.createStatement()) { statement =>
        resource(statement.executeQuery("SELECT name FROM things")) { rows =>
          require(rows.next() && rows.getString(1) == "Hello, world", "SQLite persistence failed")
        }
      }
    }
    val loopback = InetAddress.getLoopbackAddress()
    resource(new ServerSocket(0, 1, loopback)) { listener =>
      listener.setSoTimeout(5000)
      resource(new Socket(loopback, listener.getLocalPort())) { client =>
        resource(listener.accept()) { server =>
          client.setSoTimeout(5000)
          server.setSoTimeout(5000)
          for (sender, receiver) <- Seq((client, server), (server, client)) do
            val message = ujson.write(ujson.Obj("message" -> "Hello, world"))
            sender.getOutputStream().write((message + "\n").getBytes(UTF_8))
            sender.getOutputStream().flush()
            val reader = new java.io.BufferedReader(new java.io.InputStreamReader(receiver.getInputStream(), UTF_8))
            require(ujson.read(reader.readLine())("message").str == "Hello, world", "JSON over TCP failed")
        }
      }
    }
    println("Scala environment ready: SQLite persistence and JSON over TCP verified.")
  finally Files.deleteIfExists(path)
