# 参考
## - [jdk8 jstat使用](https://blog.csdn.net/maosijunzi/article/details/46049117)
## - [JVM调优浅谈](https://www.cnblogs.com/xingzc/p/5756119.html)

## jps
- jps（JVM Process Status Tool，虚拟机进程监控工具），这个命令可以列出正在运行的虚拟机进程，并显示虚拟机执行主类名称，以及这些进程的本地虚拟机唯一ID。这个ID被称为本地虚拟机唯一ID（local virtual Machine Identifier，简写为LVMID）。如果你在linux的一台服务器上使用jps得到的LVMID其实就是和ps 命令得到的PID是一样的。

## jinfo
- jinfo （Configuration Info for Java，配置信息工具）
这个命令可以实时地查看和调整虚拟机各项参数。

## jhat
- jhat（虚拟机堆转储快照分析工具），这个工具是用来分析jmap dump出来的文件。
由于这个工具功能比较简陋，运行起来也比较耗时，所以这个工具不推荐使用，推荐使用MAT。

## jmap
- jmap（Memory Map for Java，内存映像工具），用于生成堆转存的快照，一般是heapdump或者dump文件。
如果不适用jmap命令，可以使用-XX:+HeapDumpOnOutOfMemoryError参数，当虚拟机发生内存溢出的时候可以产生快照。
或者使用kill -3 pid也可以产生。jmap的作用并不仅仅是为了获取dump文件，
它可以查询finalize执行队列，java堆和永久代的详细信息，如空间使用率，当前用的哪种收集器。

## jstat
- jstat（JVM Statistics Monitoring Tool，虚拟机统计信息监视工具），这个命令用于监视虚拟机各种运行状态信息。
它可以显示本地或者远程虚拟机进程中的类装载、内存、垃圾收集、JIT编译等运行数据，虽然没有GUI图形界面，只是提供了纯文本控制台环境的服务器上，
但它是运行期间定位虚拟机性能问题的首选工具。

## jstack
- jstack（Java Stack Trace，Java堆栈跟踪工具），这个命令用于查看虚拟机当前时刻的线程快照（一般是threaddump 或者 javacore文件）。
线程快照就是当前虚拟机内每一条线程正在执行的方法堆栈的集合。生成线程快照的主要目的是定位线程出现长时间停顿的原因，
入线程间死锁、死循环、请求外部资源导致的长时间等待都是导致线程长时间停顿的常见原因。线程出现停顿的时候通过jstack来查看各个线程的调用堆栈，
就可以知道没有响应的线程到底在后台做些什么事情。

## hprof（Heap/CPU Profiling Tool）
- hprof能够展现CPU使用率，统计堆内存使用情况

## jconsole
- jconsole:一个java GUI监视工具，可以以图表化的形式显示各种数据。并可通过远程连接监视远程的服务器VM。
用java写的GUI程序，用来监控VM，并可监控远程的VM，非常易用，而且功能非常强。命令行里打 jconsole，选则进程就可以了。

## jvisualvm
jvisualvm同jconsole都是一个基于图形化界面的、可以查看本地及远程的JAVA GUI监控工具，
Jvisualvm同jconsole的使用方式一样，直接在命令行打入jvisualvm即可启动，jvisualvm界面更美观一些，数据更实时：

