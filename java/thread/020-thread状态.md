> 参考: https://juejin.im/post/5bbc9311f265da0ac6696d06#heading-11

### 线程状态：
1. NEW 尚未启动
2. RUNNABLE 正在执行中
3. BLOCKED 阻塞的（被同步锁或者IO锁阻塞）
4. WAITING 永久等待状态 
5. TIMED_WAITING 等待指定的时间重新被唤醒的状态
6. TERMINATED 执行完成

<details>
  <summary>NEW 状态</summary>
 
 ```java
Thread t = new Thread(){
    @Override
    public void run() {
        System.out.println("state new ");
    }
};

System.out.println(t.getState()); // NEW
 ```
 
</details>

<details>
  <summary>RUNNABLE 状态</summary>
 
 ```java
Thread t = new Thread(){
    @Override
    public void run() {
        System.out.println("show thread state ");
    }
};
t.start();
System.out.println(t.getState()); // RUNNABLE
 ```
 
</details>
<details>
  <summary>BLOCKED 状态</summary>
 
 ```java
 /**
 * 使用synchronized同步阻塞实现。
 */
public class ThreadStateBlokced {

    public static void main(String[] args) throws InterruptedException {
        MyCounter counter = new MyCounter();

        new Thread(() -> counter.increase()).start();

        Thread t = new Thread(()-> counter.increase());
        t.start();

        // 主线程休眠10毫秒
        Thread.sleep(10);
        // 打印线程2的状态
        System.out.println(t.getState()); // BLOCKED
    }

    static class MyCounter {
        int count = 1;

        synchronized void increase(){
            count++;

            try {
                Thread.sleep(1000 * 10);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println(Thread.currentThread().getName()+": "+count);
        }
    }
}
 ```
 
</details>
<details>
  <summary>WAITING 状态</summary>
 
 ```java
public class ThreadStateWAITING {

    public static void main(String[] args) throws InterruptedException {
        Thread t = new MyThread();
        t.start();

        Thread.sleep(10);

        System.out.println(t.getState()); // WAITING

        // 可使用notify/notifyAll方法 唤醒线程
        synchronized(MyThread.class){
//            MyThread.class.notify();
            MyThread.class.notifyAll();
        }
    }

    static class MyThread extends Thread{
        @Override
        public void run() {
            synchronized(MyThread.class){
                try {
                    MyThread.class.wait();
                    System.out.println("thread "+getName()+" after waiting....");
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
 ```
 
</details>
<details>
  <summary>TIMED_WAITING 状态</summary>
 
 ```java
 public class ThreadStateTIMED_WAITING {

    public static void main(String[] args) throws InterruptedException {
        Thread t = new MyThread();
        t.start();

        Thread.sleep(10);

        System.out.println(t.getState()); // TIMED_WAITING
    }

    static class MyThread extends Thread{
        @Override
        public void run() {
            synchronized (MyThread.class) {
                try {
                    MyThread.class.wait(2000);
                    System.out.println("after Thread "+getName()+" timed wait!");
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }

        }
    }
}
 ```
 
</details>
<details>
  <summary>TERMINATED 状态</summary>
 
 ```java
 public class ThreadStateTERMINATED {
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread(() -> System.out.println(Thread.currentThread().getName()+" do task!"));
        t.start();

        Thread.sleep(100);

        System.out.println(t.getState()); //TERMINATED
    }
}
 ```
 
</details>


<details>
  <summary>死锁</summary>
 
 ```java
 /**
 * 利用sleep方法不释放锁，复现死锁情况。
 * <p>
 *     线程1持有锁A，然后获取锁B.<br>
 *     线程2持有锁B,然后获取锁B.
 * </p>
 */
public class DeadLock {
    public static void main(String[] args) throws InterruptedException {
        Object lockA = new Object();
        Object lockB = new Object();

        Thread t1 = new Thread(){
            public void run() {
                synchronized(lockA){
                    System.out.println("thread 1 action!");
                    try {
                        sleep(1000);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                    synchronized (lockB){
                        System.out.println("thread 1 get lockB successfully!");
                    }
                    System.out.println("thread 1 finish!");
                }
            }
        };
        Thread t2 = new Thread(){
            public void run() {
                synchronized (lockB){
                    System.out.println("thread 2 action!");
                    try {
                        sleep(1000);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }

                    synchronized (lockA){
                        System.out.println("thread [22] get lockA successfully!");
                    }
                    System.out.println("thread 2 finish!");
                }
            }
        };

        t1.start();
        t2.start();

        Thread.sleep(100000);
    }
}
 ```
 
</details>
