

**意图：** 保证一个类仅有一个实例，并提供一个访问它的全局访问点。  

**主要解决：** 一个全局使用的类频繁地创建与销毁。   

**何时使用：** 当您想控制实例数目，节省系统资源的时候。  

**如何解决：** 判断系统是否已经有这个单例，如果有则返回，如果没有则创建。  

**关键代码：** 构造函数是私有的。  

**应用实例：** 1、一个党只能有一个书记。  
2、Windows 是多进程多线程的，在操作一个文件的时候，就不可避免地出现多个进程或线程同时操作一个文件的现象，所以所有文件的处理必须通过唯一的实例来进行。  
3、一些设备管理器常常设计为单例模式，比如一个电脑有两台打印机，在输出的时候就要处理不能两台打印机打印同一个文件。  

**优点：** 1、在内存里只有一个实例，减少了内存的开销，尤其是频繁的创建和销毁实例（比如管理学院首页页面缓存）。  
2、避免对资源的多重占用（比如写文件操作）。  

**缺点：** 没有接口，不能继承，与单一职责原则冲突，一个类应该只关心内部逻辑，而不关心外面怎么样来实例化。  

**使用场景：** 1、要求生产唯一序列号。   
2、WEB 中的计数器，不用每次刷新都在数据库里加一次，用单例先缓存起来。  
3、创建的一个对象需要消耗的资源过多，比如 I/O 与数据库的连接等。  

**注意事项：** getInstance() 方法中需要使用同步锁 synchronized (Singleton.class) 防止多线程同时进入造成 instance 被多次实例化。


```java
public class s03Singleton {

    public static void main(String[] args) {
        Singleton5.init();
        Singleton5 instance = Singleton5.getInstance();
    }

}

/**
 * 一。
 是否 Lazy 初始化：是
 是否多线程安全：否
 实现难度：易
 */
class Singleton1{
    private static Singleton1 instance;
    private Singleton1 (){}

    public static Singleton1 getInstance() {
        if (instance == null) {
            instance = new Singleton1();
        }
        return instance;
    }
}

/**
 * 二。
 是否 Lazy 初始化：是
 是否多线程安全：是
 实现难度：易
 */
class Singleton2{
    private static Singleton2 instance;
    private Singleton2 (){}
    public static synchronized Singleton2 getInstance() {
        if (instance == null) {
            instance = new Singleton2();
        }
        return instance;
    }
}

/**
 * 三。
 是否 Lazy 初始化：否
 是否多线程安全：是
 实现难度：易
 */
class Singleton3{
    private static Singleton3 instance = new Singleton3();
    private Singleton3 (){}
    public static Singleton3 getInstance() {
        return instance;
    }
}

/**
 * 四。
 是否 Lazy 初始化：是
 是否多线程安全：是
 实现难度：较复杂
 */
class Singleton4{
    private volatile static Singleton4 singleton;
    private Singleton4 (){}
    public static Singleton4 getSingleton() {
        if (singleton == null) {
            synchronized (Singleton4.class) {
                if (singleton == null) {
                    singleton = new Singleton4();
                }
            }
        }
        return singleton;
    }
}

/**
 * 五。
 是否 Lazy 初始化：是
 是否多线程安全：是
 实现难度：一般
 */
class Singleton5{
    private static class SingletonHolder {
        private static final Singleton5 INSTANCE = new Singleton5();
    }
    private Singleton5 (){
        System.out.println("create Singleton5");
    }
    public static final Singleton5 getInstance() {
        return SingletonHolder.INSTANCE;
    }
    
    public static void init(){
        System.out.println("init Singleton5");
    }
}

```
