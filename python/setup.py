from src.database import database
from src.server import server
from threading import Thread
import socket
import unittest


class TestSetup(unittest.TestCase):
    @classmethod
    def setUpClass(self):
        # Initialize database
        self.db = database.Database()
        self.db.initialize()

        # Initialize server on background thread
        self.server = server.Server((server.Server.DEFAULT_SERVER_HOST,
                                     server.Server.DEFAULT_SERVER_PORT), server.ServerHandler)
        Thread(target=self.server.serve_forever).start()

        # Initialize mock client to send socket messages to server
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect((server.Server.DEFAULT_SERVER_HOST,
                             server.Server.DEFAULT_SERVER_PORT))

    @classmethod
    def tearDownClass(self):
        self.socket.close()
        self.server.shutdown()
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
        self.socket.send(b'Hello, world')

        result = self.socket.recv(1024).decode()

        self.assertEqual(result, 'HELLO, WORLD')


if __name__ == "__main__":
    unittest.main()
