## 需求
允许绿化软件固定到任务栏， 快速从任务栏双击启动

## 具体步骤
1. `/usr/share/applications`下创建一个 `.desktop` 文件
2. 为我们自己的软件，自定义 `.desktop` 文件
```
[Desktop Entry]
Version=1.0
Name=Idea
Comment=java ide
GenericName=ji
Keywords=java;ide
Exec=/home/gx/software/jet-idea-IU/bin/idea.sh   # 我们应用的路径
Terminal=false
StartupWMClass=jetbrains-idea  # 应用类别名
Type=Application
Icon=/home/gx/software/jet-idea-IU/bin/idea.png
Categories=GNOME;GTK;ide
StartupNotify=true
```

## 易错点
1. `Name` 与 `.desktop` 文件名需一致
比如说，   `Name=Idea`， 那么文件名需为 `Idea.desktop`
2. `StartupWMClass` 属性填写不正确， 会导致任务栏的软件图标出错
可执行命令 `xprop WM_CLASS`，鼠标会变成一个加号， 然后点击目标应用的窗口，可获得相应类别
比如说， idea 窗口的结果如下：
```bash
gx@MX:~
$ xprop WM_CLASS
WM_CLASS(STRING) = "jetbrains-idea", "jetbrains-idea"
```
3. 应用路径可为绝对路径， 也可以是  环境变量`PATH`下可惜找到的可执行文件的相对路径