import sqlite3
import os
import time


class Database:
    DATASTORE_PATH = os.path.join(os.path.dirname(__file__), './datastore.db')
    SCHEMA_PATH = os.path.join(os.path.dirname(__file__), './schema.sql')

    def initialize(self):
        self.db = sqlite3.connect(
            self.DATASTORE_PATH, isolation_level=None, check_same_thread=False)

        # Return back a dictionary keyed by column for row values
        self.db.row_factory = sqlite3.Row

        with open(self.SCHEMA_PATH, 'r') as f:
            schema_create = f.read()
        self.db.executescript(schema_create)

        return self.db

    def write(self):
        sql = '''
        INSERT INTO messages (msg)
        VALUES (?)
        '''

        message = 'Current time: ' + str(time.time())

        return self.db.execute(sql, (message,)).lastrowid

    def close(self):
        self.db.close()

    def reconnect(self):
        self.close()
        self.initialize()

    def cleanup(self):
        self.close()
        os.remove(self.DATASTORE_PATH)
