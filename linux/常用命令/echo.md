echo是一个非常简单、直接的linux命令, 将argument送出至标准输出（STDOUT），通常就是在显示器（monitor）上输出。


### 格式
```
echo [-options] [argument]...
```

### 常用选项[-options]
```
    -e：启用反斜线控制字符的转换
    -E：关闭反斜线控制字符的转换（预设如此）
    -n：取消行末之换行符号（与 -e 选项下的 \c 字符同意）
```

echo 命令所支持的反斜线控制字符如下：
```
    \a：ALERT / BELL (从系统喇叭发送出声音)
    \b：BACKSPACE ，也就是向左刪除
    \c：取消行末的行换符号
    \E：ESCAPE，逃离键
    \f：FORMFEED，换页字符
    \n：NEWLINE，換行字符
    \r：RETURN，回车键
    \t：TAB，表格跳位鍵
    \v：VERTICAL TAB，垂直表格跳位鍵
    \n：ASCII 八进制编码(以 x 为十六进制)
    \\：反斜线本身
```

### 例子

-e 选项
```bash
$ echo -e "a\tb\tc\nd\te\tf"
a       b       c
d       e       f
```

-n 选项
```bash
$ echo first line
first line
$ echo -n first line
first line $
```

-e 与 -n 结合使用
```bash
$ echo -ne "a\tb\tc\nd\te\bf\a"
a       b       c
d       f $
```



