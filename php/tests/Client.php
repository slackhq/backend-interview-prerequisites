<?php

final class Client {
    private $socket;

    public function connect() {
        $this->socket = stream_socket_client("tcp://localhost:8000", $errno, $errstr, 5, STREAM_CLIENT_CONNECT | STREAM_CLIENT_PERSISTENT);
        if (!$this->socket) {
            throw new \Exception("Unable to connect to socket. Is the server started?");
        }
        stream_set_blocking($this->socket, 0);
        return $this->socket;
    }

    public function disconnect() {
        stream_socket_shutdown($this->socket, STREAM_SHUT_RDWR);
        fclose($this->socket);
        unset($this->socket);
    }

    public function sendRequest(string $message): string {
        stream_socket_sendto($this->socket, $message);
        return $this->listenForResponse();
    }

    public function listenForResponse(): string {
        while (true) {
            $contents = fread($this->socket, 8192);
            if ($contents) {
                return $contents;
            }
        }
    }
}