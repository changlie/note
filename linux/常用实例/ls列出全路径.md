### linux ls 列出全路径

1、列出当前目录的文件、文件夹完整路径
   ls -1 |awk '{print i$0}' i=`pwd`'/'

2、列出当前目录及子目录的文件、文件夹完整路径
   ls -R |awk '{print i$0}' i=`pwd`'/'

2b） 列出当前目录及子目录下的文件夹完整路径
   ls -FR | grep /|sed"s::‘pwd‘/:"

3、用find实现，好像运行要慢些
   find/−name"∗.∗"−execls

4、递归列出当前目录及子目录名称
   ls −FR|grep/

5、递归列出当前目录及子目录名称，包括相关属性
    ls -lR | grep "^d"
    # drwxr-xr-x 3 idea idea  4096 Aug  2  2009 images

6、只列出当前目录下的子目录
    用ls只列出子目录
    ls -d */