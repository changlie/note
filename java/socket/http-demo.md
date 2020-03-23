## 用socket写一个简单的http服务器

```java
public class Server {
    public static void main(String[] args) {
        ExecutorService threadPool = Executors.newFixedThreadPool(3);

        try(ServerSocket server = new ServerSocket(9999)){
            System.out.println("server:9999 is running!!!");
            while (true){
                Socket req = server.accept();
                threadPool.execute(new HttpHandler(req));
            }

        }catch (Exception e){
            e.printStackTrace();
        }
    }
}

class HttpHandler implements Runnable {
    private Socket req;

    public HttpHandler(Socket req) {
        this.req = req;
    }

    @Override
    public void run() {
        try {
            InputStream in = req.getInputStream();
            int len = in.available();
            byte[] buf = new byte[len];
            int read = in.read(buf, 0, len);
            if(read>0){
                System.out.println(new String(buf));
            }

            PrintStream out = new PrintStream(req.getOutputStream());

            String content = "<h1>hello world!</h1>";

            StringBuilder headers = new StringBuilder();
            headers.append("HTTP/1.1 200 OK").append(System.lineSeparator());
            headers.append("Connection: keep-alive").append(System.lineSeparator());
            headers.append("Content-Length: ").append(content.getBytes("utf-8").length).append(System.lineSeparator());
            headers.append("Pragma: no-cache").append(System.lineSeparator());
            headers.append(System.lineSeparator());

            out.print(headers.toString());
            out.println(content);

            req.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```
