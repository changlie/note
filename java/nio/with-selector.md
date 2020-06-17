client
```java
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.SocketChannel;
import java.util.Iterator;

public class NoBlockClient {
    public static void main(String[] args) throws IOException {
        SocketChannel socketChannel = SocketChannel.open(new InetSocketAddress("127.0.0.1", 7777));

        socketChannel.configureBlocking(false);

        // 1.2获取选择器
        Selector selector = Selector.open();

        // 1.3将通道注册到选择器中，获取服务端返回的数据
        socketChannel.register(selector, SelectionKey.OP_READ);

        ByteBuffer buf = ByteBuffer.allocate(1024);
        buf.put("v2 hey, nio server !".getBytes());

        buf.flip();
        socketChannel.write(buf);
        buf.clear();



        if (selector.select() > 0) {
            // 6. 获取当前选择器所有注册的“选择键”(已就绪的监听事件)
            Iterator<SelectionKey> iterator = selector.selectedKeys().iterator();

            // 7. 获取已“就绪”的事件，(不同的事件做不同的事)
            while (iterator.hasNext()) {

                SelectionKey selectionKey = iterator.next();

                // 8. 读事件就绪
                if (selectionKey.isReadable()) {

                    // 8.1得到对应的通道
                    SocketChannel channel = (SocketChannel) selectionKey.channel();

                    ByteBuffer responseBuffer = ByteBuffer.allocate(1024);

                    // 9. 知道服务端要返回响应的数据给客户端，客户端在这里接收
                    int readBytes = channel.read(responseBuffer);

                    if (readBytes > 0) {
                        // 切换读模式
                        responseBuffer.flip();
                        System.out.println(new String(responseBuffer.array(), 0, readBytes));
                    }
                }

                // 10. 取消选择键(已经处理过的事件，就应该取消掉了)
                iterator.remove();
            }
        }

//        socketChannel.close();
    }
}

```

server
```java
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.*;
import java.util.Iterator;

public class NoBlockServer {
    public static void main(String[] args) throws IOException {
        ServerSocketChannel server = ServerSocketChannel.open();
        server.configureBlocking(false);
        server.bind(new InetSocketAddress(7777));

        Selector selector = Selector.open();
        server.register(selector, SelectionKey.OP_ACCEPT);

        while(selector.select() > 0){
            Iterator<SelectionKey> iterator = selector.selectedKeys().iterator();

            while(iterator.hasNext()) {
                SelectionKey key = iterator.next();
                if (key.isAcceptable()) {
                    SocketChannel client = server.accept();
                    client.configureBlocking(false);
                    client.register(selector, SelectionKey.OP_READ);
                }else if (key.isReadable()) {
                    SocketChannel client = (SocketChannel) key.channel();
                    try{
                        ByteBuffer buf = ByteBuffer.allocate(1024);
                        int len = 0 ;
                        while ((len = client.read(buf))>0) {
                            buf.flip();
                            System.out.println(new String(buf.array(), 0, len));
                            buf.clear();
                        }

                        buf.put(("hi, nio boy No."+System.currentTimeMillis()).getBytes());
                        buf.flip();
                        client.write(buf);
                        buf.clear();
                    }catch (Exception e){
//                        e.printStackTrace();
                        System.out.println(e.getMessage());
                        client.finishConnect();
                        client.close();
                    }
                }
                iterator.remove();
            }
        }

        System.out.println("server is over!");
    }
}
```
