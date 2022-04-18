from src.database import database
from src.server import server
from src.client import client
import unittest


class TestSetup(unittest.TestCase):
    @classmethod
    def setUpClass(self):
        # Initialize database
        self.db = database.Database()
        self.db.initialize()

        # Initialize server on background thread
        self.server = server.Server()
        self.server.initialize()

        # Initialize mock client to send socket messages to server
        self.client = client.Client(
            server_host=server.Server.DEFAULT_SERVER_HOST, server_port=server.Server.DEFAULT_SERVER_PORT)
        self.client.initialize()

    @classmethod
    def tearDownClass(self):
        self.client.cleanup()
        self.server.cleanup()
        self.db.cleanup()

    def test_db(self):
        # First insert should have id 1
        id = self.db.write()
        self.assertEqual(id, 1)

        # Close database connection and reconnect
        self.db.reconnect()

        # Second insert should have id 2
        id = self.db.write()
        self.assertEqual(id, 2)

    def test_server(self):
        result = self.client.send('Hello, world')

        self.assertEqual(result, 'HELLO, WORLD')


if __name__ == "__main__":
    unittest.main()
