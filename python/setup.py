from src.database import database
from src.server import server
from threading import Thread
import socket


def send_message():
    print('hello')
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.connect((server.Server.DEFAULT_SERVER_HOST,
                  server.Server.DEFAULT_SERVER_PORT))
    sock.sendall(b'Hello, world')
    data = sock.recv(1024)
    sock.close()
    print('Received', repr(data))


if __name__ == "__main__":
    db = database.Database()

    db.initialize()
    db.write()
    db.close()
    db.initialize()
    db.write()
    db.cleanup()

    s = server.Server((server.Server.DEFAULT_SERVER_HOST,
                      server.Server.DEFAULT_SERVER_PORT), server.ServerHandler)

    t1 = Thread(target=s.serve_forever)
    t2 = Thread(target=send_message)

    t1.start()
    t2.start()

    t2.join()

    s.shutdown()
