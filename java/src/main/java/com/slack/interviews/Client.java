package com.slack.interviews;

import java.io.IOException;
import java.net.InetAddress;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SocketChannel;

public class Client {

    private SocketChannel socketChannel;

    public Client(String hostname, int port) throws IOException {
        this.socketChannel = SocketChannel.open(new InetSocketAddress(InetAddress.getByName(hostname), port));
    }

    public void write(String message) throws IOException {
        ByteBuffer buffer = ByteBuffer.wrap(message.getBytes());
        socketChannel.write(buffer);
    }

    public String read() throws IOException {
        ByteBuffer byteBuffer = ByteBuffer.allocate(1024);
        socketChannel.read(byteBuffer);
        return new String(byteBuffer.array()).trim();
    }

    public void close() throws IOException {
        socketChannel.close();
    }
}