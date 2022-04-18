import socket


class Client():
    def __init__(self, server_host, server_port):
        self.server_host = server_host
        self.server_port = server_port

    def initialize(self):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect((self.server_host, self.server_port))

    def cleanup(self):
        self.socket.close()

    def send(self, message):
        self.socket.send(bytes(message, encoding='utf8'))
        return self.socket.recv(1024).decode()
