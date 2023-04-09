环境： ubuntu20 
gnome版本：3.36.8


### 系统壁纸路径
/usr/share/backgrounds


### Gnome2
用以下命令可查看当前背景
```bash
gconftool-2 --get /desktop/gnome/background/picture_filename
```

用以下命令可设置当前背景
```bash
gconftool-2 --set --type string /desktop/gnome/background/picture_filename '/home/feichashao/Desktop/feichashao.jpg'
```


### Gnome3
用以下命令可查看当前背景
```bash
gsettings get org.gnome.desktop.background picture-uri
```
用以下命令可修改当前背景
```bash
gsettings set org.gnome.desktop.background picture-uri 'file:///home/feichashao/Desktop/feichashao.jpg'
```

> 注：
1. 如果通过ssh执行上面的设置命令，会出现报错。
2. 要更改用户A的背景，应该用A用户执行上述命令。
```
(process:10123): dconf-WARNING **: failed to commit changes to dconf: Error spawning command line `dbus-launch --autolaunch=d954c83dff5ff4d8570c01048fa36b50 --binary-syntax --close-stderr': Child process exited with code 1
```












