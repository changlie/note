### 简单的TCP实现

#### 服务端
```java
public class Server {
    public static void main(String[] args) {
        ExecutorService threadPool = Executors.newFixedThreadPool(3);

        try(ServerSocket server = new ServerSocket(9999)){
            System.out.println("server:9999 is running!!!");
            while (true){
                Socket req = server.accept();
                threadPool.execute(new RequestHandler(req));
            }

        }catch (Exception e){
            e.printStackTrace();
        }
    }
}

class RequestHandler implements Runnable {
    private Socket req;

    public RequestHandler(Socket req) {
        this.req = req;
    }

    @Override
    public void run() {
        try {
            InputStream in = req.getInputStream();
            BufferedReader r = new BufferedReader(new InputStreamReader(in, "utf8"));
            OutputStream out = req.getOutputStream();
            PrintStream pout = new PrintStream(out);
            String tmp;
            while ((tmp=r.readLine())!=null){
                if("bye".equals(tmp)){
                    break;
                }
                System.out.println("client: "+tmp);
                pout.println(tmp.toUpperCase());
            }
            req.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

```

#### 客户端
```java
public class Client {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        try(Socket client = new Socket("127.0.0.1", 9999)){
            BufferedReader r = new BufferedReader(new InputStreamReader(client.getInputStream()));
            PrintStream out = new PrintStream(client.getOutputStream());

            String s;
            while(!"bye".equals(s = scanner.nextLine())){
                out.println(s);
                System.out.println("Server: "+r.readLine());
            }

        }catch (Exception e){
            e.printStackTrace();
        }
        scanner.close();
    }
}

```
