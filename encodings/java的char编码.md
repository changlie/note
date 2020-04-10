## 转载

### 问题: Java 语言中一个字符占几个字节？

### 答：脱离具体的编码谈某个字符占几个字节是没有意义的。
如果“字符”是指 Java 中的 char，那好，那它就是 16 位，2 字节。<br>
如果“字符”是指我们用眼睛看到的那些“抽象的字符”，那么，谈论它占几个字节是没有意义的。
```
就好比有一个抽象的整数“42”，你说它占几个字节？这得具体看你是用 byte，short，int，还是 long 来存它。
用 byte 存就占一字节，用 short 存就占两字节，int 通常是四字节，long 通常八字节。
当然，如果你用 byte，受限于它有限的位数，有些数它是存不了的，比如 256 就无法放在一个 byte 里了。
```
#### 1. 同一个字符在不同的编码下可能占不同的字节。
```
举“字”字为例.
在 GBK 编码下占 2 字节，
在 UTF-16 编码下也占 2 字节，
在 UTF-8 编码下占 3 字节，
在 UTF-32 编码下占 4 字节。
```

#### 2. 不同的字符在同一个编码下也可能占不同的字节。
```
“字”在 UTF-8 编码下占3字节，而“A”在 UTF-8 编码下占 1 字节。（因为 UTF-8 是变长编码）
```

而 Java 中的 char 本质上是 UTF-16 编码。而 UTF-16 实际上也是一个变长编码（2 字节或 4字节）。

如果一个抽象的字符在 UTF-16 编码下占 4 字节，显然它是不能放到 char 中的。换言之， char 中只能放 UTF-16 编码下只占 2 字节的那些字符。

而 getBytes 实际是做编码转换，你应该显式传入一个参数来指定编码，否则它会使用缺省编码来转换。
```
你说“ new String(“字”).getBytes().length 返回的是3 ”，这说明缺省编码是 UTF-8.
如果你显式地传入一个参数，比如这样“ new String(“字”).getBytes(“GBK“).length ”，那么返回就是 2.
```
你可以在启动 JVM 时设置一个缺省编码，
```
假设你的类叫 Main，那么在命令行中用 java 执行这个类时可以通过 file.encoding 参数设置一个缺省编码。
比如这样：java -Dfile.encoding=GBK Main
这时，你再执行不带参数的 getBytes() 方法时，new String(“字”).getBytes().length 返回的就是 2 了，因为现在缺省编码变成 GBK 了。
当然，如果这时你显式地指定编码，new String(“字”).getBytes(“UTF-8“).length 返回的则依旧是 3.
```
否则，会使用所在操作系统环境下的缺省编码。
```
通常，Windows 系统下是 GBK，Linux 和 Mac 是 UTF-8.
但有一点要注意，在 Windows 下使用 IDE 来运行时，比如 Eclipse，如果你的工程的缺省编码是 UTF-8，在 IDE 中运行你的程序时，会加上上述的 -Dfile.encoding=UTF-8 参数，这时，即便你在 Windows 下，缺省编码也是 UTF-8，而不是 GBK。
```
由于受启动参数及所在操作系统环境的影响，不带参数的 getBytes 方法通常是不建议使用的，最好是显式地指定参数以此获得稳定的预期行为。


## 例子
java char 默认使用utf-16进行编码
```java
public class CharTest1 {
    public static void main(String[] args) {
        // 常见的四字节字符：𪨊𫞄𠀀𫭟𫭢𫭼
        String s = "天上𥊍𪚥";
        System.out.println("len: "+ s.length());
        for(int i=0; i<s.length(); i++){
            System.out.println(i+": "+s.charAt(i));
        }
        System.out.println("source: "+ s);
    }
}
```
输出:
```
len: 6
0: 天
1: 上
2: ?
3: ?
4: ?
5: ?
source: 天上𥊍𪚥
```

> 参考:

[Java 语言中一个字符占几个字节](https://xiaogd.net/java-%E8%AF%AD%E8%A8%80%E4%B8%AD%E4%B8%80%E4%B8%AA%E5%AD%97%E7%AC%A6%E5%8D%A0%E5%87%A0%E4%B8%AA%E5%AD%97%E8%8A%82%EF%BC%9F/)

[4字节UTF16编码和对C++/Java的影响](https://blog.csdn.net/jiangfuqiang/article/details/17222255)

