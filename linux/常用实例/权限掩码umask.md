> 转: [Linux下的权限掩码umask](https://www.jianshu.com/p/a092bb562a4c)

权限掩码是用于指定文件,目录创建时的权限默认值的
> Linux下的etc/profile和etc/bashrc中都有默认的umask设置    
`umask`可以查看与设置权限掩码   
查看权限掩码
```
umask
```
查看目录创建时的权限默认值, 去掉执行权限`x`就文件创建时的默认权限值了, 因为文件创建时是没有执行权限的
```
umask -S
```
设置权限掩码
```
umask 0022
```
> 0666 - 权限掩码  = 文件权限默认值  
0777 - 权限掩码  = 目录权限默认值

### 实例
```bash
changlie@as:~
$ mkdir test; cd test
changlie@as:~/test
$ umask
0002
changlie@as:~/test
$ umask -S
u=rwx,g=rwx,o=rx
changlie@as:~/test
$ touch f1; mkdir dir1
changlie@as:~/test
$ ls -l
总用量 4
drwxrwxr-x 2 changlie changlie 4096 6月  13 12:35 dir1
-rw-rw-r-- 1 changlie changlie    0 6月  13 12:35 f1
```
重新设置权限掩码, 再试
```bash
changlie@as:~/test
$ umask 0077
changlie@as:~/test
$ umask -S
u=rwx,g=,o=
changlie@as:~/test
$ touch f2; mkdir dir2
changlie@as:~/test
$ ls -l
drwx------ 2 changlie changlie 4096 6月  13 12:38 dir2
-rw------- 1 changlie changlie    0 6月  13 12:38 f2
```

### 号外
在目录和文件的默认权限属性是不同的，因为对于一个目录来说它的x权限也就是执行权限是很重要的，进入目录等操作都是需要目录具有执行权限的，而对于文件来说，一般情况都是用于数据的记录操作，所以一般不需要执行权限。从而，在linux下默认的情况是这样的：    
- 如果用户创建的是目录，则默认所有权限都开放，为777，默认为：drwxrwxrwx   
- 如果创建的是文件，默认没有x权限，那么就只有r、w两项，最大值为666，默认为：-rw-rw-rw-   
000 - 777(对目录)  000 - 666（对文件）   
r、w、x分别是4、2、1，要拿掉读权限就输入4，拿掉写权限就输入2，以此类推。
