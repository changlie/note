synchronized提供内部锁的机制，防止其它线程同时进入synchronized的代码块。synchronized由两部分组成：1.锁对象的引用；2.锁保护的代码块。

对锁对象引用的不同，是static synchronized和synchronized最大的区别：  
```java
public class SynchronizedTester {
    /*
     * testOne()和testTwo()均是对SynchronizedTester这个类的对象(实例)加上锁
     * 也就是锁对象的引用是SynchronizedTester这个类的对象(实例)
     * 两种写法基本等价
     */
    private synchronized void testOne(){
        //do something...
    }
    private void testTwo(){
        synchronized(this){
            //do something...
        }
    }
}
```
```java
public class SynchronizedStaticTester {    
    /*
     * testOne()和testTwo()则是对SynchronizedTester这个类加上锁
     * 也就是锁对象的引用是SynchronizedTester这个类，不再是类的对象或类的实例
     * 两种写法基本等价
     */
    private static synchronized void testOne(){
        //do sth.
    }
    private static void testTwo(){
        synchronized (SynchronizedStaticTester.class) { //Notice: Not this!
            // do sth.
        }
    }
}
```
## synchronized的范围是某个类的对象/实例，防止多个线程同时访问同一个类对象/实例的synchronized代码块。
## static synchronized地方范围是某个类，防止多个线程同时访问这个类的synchronized代码块。
