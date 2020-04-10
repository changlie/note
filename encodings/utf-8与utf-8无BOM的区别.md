BOM——Byte Order Mark，就是字节序标记

BOM签名的意思就是告诉编辑器当前文件采用何种编码,是小字节序还是大字节序

BOM对UFT-8没有作用,是为了兼容UTF-16,UTF-32才加上的

所以说，utf-8无BOM才是真正的utf8


UTF-8不需要BOM来表明字节顺序，但可以用BOM来表明编码方式。
字符"ZERO WIDTH NO-BREAK SPACE"的UTF-8编码是EF BB BF。
所以如果接收者收到以EF BB BF开头的字节流，就知道这是UTF-8编码了。

在utf-8编码文件中BOM在文件头部，占用三个字节，用来标识该文件属于utf-8编码，现在已经有很多软件识别BOM头，但还是有些不能识别BOM头，比如PHP就不能识别BOM头，这也就是用记事本编辑utf-8编码的PHP文件后，就会报错的原因。


### 参考

[UTF-8(无BOM）是什么意思](https://zhidao.baidu.com/question/59267263.html)

[UTF8 与 UTF8 +BOM 区别](https://my.oschina.net/JKOPERA/blog/309423)
