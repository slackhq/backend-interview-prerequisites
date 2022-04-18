import socketserver
from threading import Thread


class Server(socketserver.ThreadingMixIn, socketserver.TCPServer):
    DEFAULT_SERVER_HOST = 'localhost'
    DEFAULT_SERVER_PORT = 8000
    allow_reuse_address = True

    def __init__(self):
        super().__init__((self.DEFAULT_SERVER_HOST,
                          self.DEFAULT_SERVER_PORT), ServerHandler)

    def initialize(self):
        Thread(target=self.serve_forever).start()

    def cleanup(self):
        self.shutdown()


class ServerHandler(socketserver.BaseRequestHandler):
    def handle(self):
        # self.request is the TCP socket connected to the client
        self.data = self.request.recv(1024).strip()
        # just send back the same data, but upper-cased
        self.request.sendall(self.data.upper())
