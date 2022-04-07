package com.slack.interviews;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

import java.io.IOException;
import java.sql.SQLException;

import java.util.Optional;

import org.json.JSONObject;
import org.junit.Test;

public class AppTest {

    public static final String HOSTNAME = "localhost";
    public static final int PORT = 8000;
    public static final String FILESYSTEM_DATABASE_URL = "jdbc:sqlite:./datastore.db";

    @Test
    public void testSockets() throws IOException {

        Server server = new Server(HOSTNAME, PORT);
        server.listenAsync();

        Client client = new Client(HOSTNAME, PORT);
        client.write("Hello");

        String response = client.read();

        assertEquals("Hello", response);

        client.close();
        server.stop();
    }

    @Test
    public void testDatabase() throws ClassNotFoundException, SQLException, IOException {
        Database database = new Database(FILESYSTEM_DATABASE_URL);
        Optional<Long> id = database.addRow("Test text");

        assertTrue(id.isPresent());

        Optional<String> text = database.getById(id.get());
        assertTrue(text.isPresent());

        assertEquals("Test text", text.get());
    }

    @Test
    public void testJson() {
        Entity entity = Entity.fromJSON((new JSONObject())
                .put("string", "Test string")
                .put("integer", 42));
        assertEquals("Test string", entity.getString());
        assertEquals(42, entity.getInteger());
    }
}
