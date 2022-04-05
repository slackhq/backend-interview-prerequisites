<?php

namespace BackendInterviewPrerequisites;

final class EchoServer {
    public function listen() {
        $socket = \stream_socket_server('tcp://localhost:8000', $errno, $errstr, \STREAM_SERVER_BIND | \STREAM_SERVER_LISTEN);
        \stream_set_blocking($socket, 0);
        $conn = \stream_socket_accept($socket);
        if ($conn) {
            while (true) {
                $contents = \fread($conn, 8192);
                if ($contents) {
                    \fwrite($conn, $contents);
                }
            }
        }
    }
}