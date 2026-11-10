import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import java.nio.file.Files
import java.sql.DriverManager

class Setup {
    static void main(String[] args) {
        def path = Files.createTempFile('interview-groovy-', '.db')
        try {
            DriverManager.getConnection("jdbc:sqlite:$path").withCloseable { db ->
                db.createStatement().withCloseable { statement ->
                    statement.executeUpdate('CREATE TABLE things (name TEXT)')
                    statement.executeUpdate("INSERT INTO things VALUES ('Hello, world')")
                }
            }
            DriverManager.getConnection("jdbc:sqlite:$path").withCloseable { db ->
                db.createStatement().withCloseable { statement ->
                    statement.executeQuery('SELECT name FROM things').withCloseable { rows ->
                        assert rows.next() && rows.getString(1) == 'Hello, world' : 'SQLite persistence failed'
                    }
                }
            }
            def loopback = InetAddress.loopbackAddress
            new ServerSocket(0, 1, loopback).withCloseable { listener ->
                listener.soTimeout = 5000
                new Socket(loopback, listener.localPort).withCloseable { client ->
                    listener.accept().withCloseable { server ->
                        client.soTimeout = 5000
                        server.soTimeout = 5000
                        [[client, server], [server, client]].each { pair ->
                            def message = JsonOutput.toJson([message: 'Hello, world'])
                            pair[0].outputStream.write((message + '\n').getBytes('UTF-8'))
                            pair[0].outputStream.flush()
                            def received = pair[1].inputStream.newReader('UTF-8').readLine()
                            assert new JsonSlurper().parseText(received).message == 'Hello, world' : 'JSON over TCP failed'
                        }
                    }
                }
            }
            println 'Groovy environment ready: SQLite persistence and JSON over TCP verified.'
        } finally {
            Files.deleteIfExists(path)
        }
    }
}
