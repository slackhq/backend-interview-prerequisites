package com.slack.interviews;

import java.io.IOException;
import java.net.InetAddress;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;
import java.sql.SQLException;

import java.util.Iterator;
import java.util.concurrent.atomic.AtomicBoolean;

public class Server {

    private ServerSocketChannel serverSocketChannel;
    private Selector selector;

    public AtomicBoolean isRunning = new AtomicBoolean(false);

    public Server(String hostname, int port) throws IOException {

        Selector selector = Selector.open();
        ServerSocketChannel serverSocketChannel = ServerSocketChannel.open();
        serverSocketChannel.configureBlocking(false);
        serverSocketChannel.bind(
                new InetSocketAddress(InetAddress.getByName(hostname), port));
        serverSocketChannel.register(selector, SelectionKey.OP_ACCEPT);
        this.serverSocketChannel = serverSocketChannel;
        this.selector = selector;
    }

    public void listenAsync() {
        new Thread(() -> {
            try {
                isRunning.set(true);
                listen();
            } catch (Throwable e) {
                System.out.println(e);
                stop();
            }
        }).start();
    }

    public void stop() {
        isRunning.set(false);
    }

    public void listen() throws IOException, ClassNotFoundException, SQLException {
        SelectionKey key = null;

        while (isRunning.get()) {
            if (selector.select() <= 0) {
                continue;
            }

            Iterator<SelectionKey> iterator = selector.selectedKeys().iterator();
            while (iterator.hasNext()) {
                key = (SelectionKey) iterator.next();
                iterator.remove();

                if (key.isAcceptable()) {
                    SocketChannel socketChannel = serverSocketChannel.accept();
                    socketChannel.configureBlocking(false);
                    socketChannel.register(selector, SelectionKey.OP_READ);
                }

                if (key.isReadable()) {
                    SocketChannel socketChannel = (SocketChannel) key.channel();
                    ByteBuffer byteBuffer = ByteBuffer.allocate(1024);
                    socketChannel.read(byteBuffer);

                    String message = new String(byteBuffer.array()).trim();

                    if (message.length() <= 0) {
                        socketChannel.close();
                        break;
                    }

                    socketChannel.write(ByteBuffer.wrap(message.getBytes()));
                }
            }
        }
    }
}
