> 转: [Bash 脚本 set 命令教程](http://www.ruanyifeng.com/blog/2017/11/bash-set.html)

### set -u
脚本在头部加上它，遇到不存在的变量就会报错，并停止执行。

### set -x
用来在运行结果之前，先输出执行的那一行命令。

### set -e
使得脚本只要发生错误，就终止执行。

### set -o pipefail
只要一个子命令失败，整个管道命令就失败，脚本就会终止执行。
