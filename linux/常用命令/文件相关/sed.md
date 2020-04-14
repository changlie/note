sed 是一种在线编辑器，它一次处理一行内容。

处理时，把当前处理的行存储在临时缓冲区中，称为“模式空间”（pattern space），接着用sed命令处理缓冲区中的内容，处理完成后，把缓冲区的内容送往屏幕。

接着处理下一行，这样不断重复，直到文件末尾。文件内容并没有 改变，除非你使用重定向存储输出, 或者使用`-i`选项

Sed主要用来自动编辑一个或多个文件；简化对文件的反复操作；编写转换程序等。

### 语法
```
sed [OPTION]... {script-only-if-no-other-script} [input-file]...
```
参数说明：
```
  -e<script>或--expression=<script> 以选项中指定的script来处理输入的文本文件。
  -f<script文件>或--file=<script文件> 以选项中指定的script文件来处理输入的文本文件。
  -h或--help 显示帮助。
  -n或--quiet或--silent 仅显示script处理后的结果。
  -V或--version 显示版本信息。
```
动作说明：
```
  a ：新增， a 的后面可以接字串，而这些字串会在新的一行出现(目前的下一行)～
  c ：取代， c 的后面可以接字串，这些字串可以取代 n1,n2 之间的行！
  d ：删除，因为是删除啊，所以 d 后面通常不接任何咚咚；
  i ：插入， i 的后面可以接字串，而这些字串会在新的一行出现(目前的上一行)；
  p ：打印，亦即将某个选择的数据印出。通常 p 会与参数 sed -n 一起运行～
  s ：取代，可以直接进行取代的工作哩！通常这个 s 的动作可以搭配正规表示法！例如 1,20s/old/new/g 就是啦！
```




> 参考

[sed命令](https://man.linuxde.net/sed#%E4%BB%8E%E6%96%87%E4%BB%B6%E8%AF%BB%E5%85%A5%EF%BC%9Ar%E5%91%BD%E4%BB%A4)

[sed 命令|菜鸟教程](https://www.runoob.com/linux/linux-comm-sed.html)
