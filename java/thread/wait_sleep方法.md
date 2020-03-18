```java
/**
 *   测试wait和sleep释放锁的测试类
 *   <p>
 *       sleep为线程的方法，而wait为Object的方法，他们的功能相似，最大本质的区别是：sleep不释放锁，wait释放锁。
 *   </p>
 *   <p>
 *       用法上的不同：sleep(milliseconds)可以用时间指定来使他自动醒过来，
 *       <br>如果时间不到你只能调用interreput()来终止线程；wait()可以用notify()/notifyAll()直接唤起。
 *   </p>
 *
 */
public class sleepWaitTest {
    public static void main(String[] args) throws Exception {
        action("wait");
//        action("sleep");
        //    执行sleep(1000)运行的结果是：1001
        //    执行wait(1000)运行的结果是：1100

        //    总结： 使用 sleep(1000)不释放同步锁，执行的是10*100+1=1001，
        //    wait(1000)释放了锁，执行的顺序是(10+1)x100=1100，所以sleep不释放锁，wait释放锁。
    }

    private static void action(String type) throws Exception {
        SynchronizedTest synchronizedTest = new SynchronizedTest();
        synchronizedTest.start();
        synchronizedTest.secord(type);
        // 主线程稍等10毫秒
        Thread.sleep(10);
        System.out.println(synchronizedTest.number);
    }

    static class SynchronizedTest extends Thread {
        int number = 10;
        public synchronized void first(){
            System.out.println("this is first!");
            number = number+1;
        }
        public synchronized void secord(String type) throws InterruptedException {
            System.out.println("this is secord!! in main thread!");
            if("wait".equals(type)){
                this.wait(1000);
            }else{
                Thread.sleep(1000);
            }
            number = number*100;
        }
        @Override
        public void run() {
            first();
        }
    }
}
```
