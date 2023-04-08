### 一、开机默认进入命令行模式

```bash
# 设置开机进入命令行
sudo systemctl set-default multi-user.target
# 重启
reboot
```
要进入图形界面，只需要输入命令startx
从图形界面切换回命令行：ctrl+alt+F7

### 二、开机默认进入图形用户界面

```bash
# 设置开机进入命令行
sudo systemctl set-default graphical.target
# 重启
reboot
```
要进入命令行模式：ctrl+alt+F2
从命令行切换到图形界面：ctrl+alt+F7