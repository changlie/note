client
```java
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.SocketChannel;

public class BlockClient  {
    public static void main(String[] args) throws IOException {
        SocketChannel socketChannel = SocketChannel.open(new InetSocketAddress("127.0.0.1", 6666));
        ByteBuffer buffer = ByteBuffer.allocate(1024);
        buffer.put("hey, block server!".getBytes());

        buffer.flip();
        socketChannel.write(buffer);
        buffer.clear();

        socketChannel.shutdownOutput();

        int len = 0;
        while ((len = socketChannel.read(buffer)) >0){
            buffer.flip();
            System.out.println(new String(buffer.array(), 0, len));
            buffer.clear();
        }

        socketChannel.close();
    }
}

```

server
```java
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.ServerSocketChannel;
import java.nio.channels.SocketChannel;

public class BlockServer  {

    public static void main(String[] args) throws IOException {
        ServerSocketChannel server = ServerSocketChannel.open();
        server.bind(new InetSocketAddress(6666));

        ByteBuffer buffer = ByteBuffer.allocate(1024);

        SocketChannel client = server.accept();

        StringBuilder ss = new StringBuilder();
        int len = -1;
        while((len = client.read(buffer)) > 0){
            buffer.flip();
            for (int i=0; i<len; i++){
                ss.append((char)buffer.get());
            }
            buffer.clear();
        }

        System.out.println("from: "+ss);

        buffer.put("it's ok".getBytes());
        buffer.flip();
        client.write(buffer);
        buffer.clear();

        client.close();
        server.close();
    }
}
```
