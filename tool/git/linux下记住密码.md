>> 参考: https://www.jianshu.com/p/29aa10ce19c1
>>> https://www.cnblogs.com/liangweiping/p/10277604.html
>>>> https://blog.csdn.net/yaorongke/article/details/88208780

1. 在终端输入
```shell
vim ~/.gitconfig
```


2.在文件输入
```
[credential]
    helper = store
```

然后再次输入密码时,密码会保存在`~/.git-credentials`文件中.
格式如下:
http://username:password@gitlab.xxx.com.cn
也可以在`~/.git-credentials`文件中直接添加.
然后再运行命令:
```
git config --global credential.helper store
```
