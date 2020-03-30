> 参考： https://www.jianshu.com/p/360e37539266
>> https://blog.csdn.net/qq496013218/article/details/69397380

应用场景： 多线程下载，断点续传

### 简介：
```
Java中的RandomAccessFile提供了对文件的读写功能。
RandomAccessFile 虽然属于java.io下的类，但它不是InputStream或者OutputStream的子类；
它也不同于FileInputStream和FileOutputStream。 
FileInputStream 只能对文件进行读操作，而FileOutputStream 只能对文件进行写操作；
与输入流和输出流不同之处就是RandomAccessFile可以访问文件的任意地方同时支持文件的读和写，并且它支持随机访问。
RandomAccessFile包含InputStream的三个read方法，也包含OutputStream的三个write方法。
同时RandomAccessFile还包含一系列的readXxx和writeXxx方法完成输入输出。
   RandomAccessFile父类：java.lang.Object
   所有接口实现：Closeable, DataInput, DataOutput, AutoCloseable
```

### 构造函数：
```
// 1.创建随机访问文件流，以从File参数指定的文件中读取，并可选择写入文件。
RandomAccessFile(File file, String mode) 
// 2.创建随机访问文件流，以从中指定名称的文件读取，并可选择写入文件。
RandomAccessFile(String name, String mode) 
```

### 构造函数中mode参数传值介绍:
```
r 代表以只读方式打开指定文件 。
rw 以读写方式打开指定文件 。
rws 读写方式打开，并对内容或元数据都同步写入底层存储设备 。
rwd 读写方式打开，对文件内容的更新同步更新至底层存储设备 。
```
  
### RandomAccessFile的重要方法：

RandomAccessFile包含了一个对象记录的指针，用于标识当前流的读写位置RandomAccessFile包含两个方法来操作文件记录指针。文件指针可以通过getFilePointer方法读取，并由seek方法设置。
 ```
// 设置文件指针偏移，从该文件的开头测量，发生下一次读取或写入。
//通俗理解：返回文件记录指针的当前位置,不指定指针的位置默认是0
 long getFilePoint()
// 设置文件指针偏移，从该文件的开头测量，发生下一次读取或写入。
//通俗理解：将文件记录指针定位到pos位置
 void seek(long pos) 
```
