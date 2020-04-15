mv命令是"move"单词的缩写，其功能大致和英文含义一样，对文件或目录进行移动或重命名

这是一个使用频率超高的文件管理命令，我们需要特别留意它与复制的区别：mv与cp的结果不同。
mv命令好像文件“搬家”，文件名称发生改变，但个数并未增加。而cp命令是对文件进行复制操作，文件个数是有增加的。

Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.

### 格式
```
       mv [OPTION]... [-T] SOURCE DEST
  or:  mv [OPTION]... SOURCE... DIRECTORY
  or:  mv [OPTION]... -t DIRECTORY SOURCE...
```
### 选项
```
  -b                           like --backup but does not accept an argument
  -f, --force                  do not prompt before overwriting
  -i, --interactive            prompt before overwrite
  -n, --no-clobber             do not overwrite an existing file
  -t, --target-directory=DIRECTORY  move all SOURCE arguments into DIRECTORY
  -T, --no-target-directory    treat DEST as a normal file(把目录当作一个普通文件)
  -u, --update                 move only when the SOURCE file is newer
                                 than the destination file or when the
                                 destination file is missing
```

### 例子
将文件file_1重命名为file_2：
```bash
[root@linuxcool ~]# mv file_1 file_2
```

将文件file移动到目录dir中 ：
```bash
[root@linuxcool ~]# mv file /dir
```

将目录dir1移动目录dir2中（前提是目录dir2已存在，若不存在则改名)：
```bash
[root@linuxcool ~]# mv /dir1 /dir2
```

将目录dir1下的文件移动到当前目录下：
```bash
[root@linuxcool ~]# mv /dir1/* .
```
