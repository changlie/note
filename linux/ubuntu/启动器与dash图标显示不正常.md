> 参考: [ubuntu启动器和dash里应用图标不正常](https://www.cnblogs.com/kylinfish/p/5044624.html)

### 问题描述
atom在启动器中的图标显示不正常,显示成了一个问号
(启动器就是下方横栏, Windows叫任务栏)

### 解决办法
ubuntu中, 启动器与dash上的应用信息是记录在以下两个目录中
```
/usr/share/applications
~/.local/share/applications
```
在这两个目录中找到atom应用的`.desktop`文件  
把`Icon`配置正确即可  
`atom.desktop`配置前
```
[Desktop Entry]
Encoding=UTF-8
Version=1.0
Type=Application
Name=Welcome Guide — Atom
Icon=atom
Path=/home/changlie/software/atom
Exec=/home/changlie/software/atom/atom
StartupNotify=false
StartupWMClass=Atom
OnlyShowIn=Unity;
X-UnityGenerated=true
```

`atom.desktop`配置后
```
[Desktop Entry]
Encoding=UTF-8
Version=1.0
Type=Application
Name=Welcome Guide — Atom
Icon=/home/changlie/software/atom/atom.png
Path=/home/changlie/software/atom
Exec=/home/changlie/software/atom/atom
StartupNotify=false
StartupWMClass=Atom
OnlyShowIn=Unity;
X-UnityGenerated=true
```
最后重启即可
