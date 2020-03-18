> 参考： https://juejin.im/post/5bbc9311f265da0ac6696d06#heading-6

<details>
<summary>继承runnable</summary>

```java
public class NewThreadByRunnable {

    public static void main(String[] args) {
        Thread t = new Thread(new MyThread());
        t.start();
    }

    static class MyThread implements Runnable{

        @Override
        public void run() {
            System.out.println("Create thread by implement interface Runnable: "+Thread.currentThread().getName());
        }
    }
}
```

</details>


<details>
<summary>继承thread</summary>

```java
public class NewThreadByExtendThread {

    public static void main(String[] args) {
        Thread t = new MyThread();
        t.start();
    }

    static class MyThread extends Thread{
        public MyThread() {
            setName("MyThread");
        }

        @Override
        public void run() {
            System.out.println("Create thread by extend class Thread!! #"+Thread.currentThread().getName());
        }
    }
}
```

</details>


<details>
<summary>继续Callable,并结合FutureTask使用，是runnable的一种补充，使线程执行后可以返回一个结果值</summary>

```java
public class NewThreadByCallable {

    public static void main(String[] args) throws ExecutionException, InterruptedException {
        FutureTask<String> f = new FutureTask<String>(new MyThread());
        Thread t = new Thread(f, "txti");
        t.start();
        System.out.println(f.get());
    }

    static class MyThread implements Callable<String> {

        @Override
        public String call() throws Exception {
            System.out.println("Create Thread by Implements interface Callable!"+Thread.currentThread().getName());
            return "日照香炉生紫烟";
        }
    }
}
```

</details>
