### linux 相关


Give root password for maintenance 解决方法

2.6 修改密码

passwd root

输入两次密码。如果提示 password updated successfully 测修改成功

2.7 输出reboot重新启动机器

2.8 以后使用新修改的密码登录即可（以下不在截图演示）

三、 Give root password for maintenance 解决方法

3.1 重新启动系统，重复步骤2.1 、2.2

3.2 按键盘上下方向键，选择 Ubuntu, with Linux 4.4.0-131-generic (recovery mode)，按下字母 e , 进入编辑命令模式

3.3 找到linux /vmlinuz-4.4.0-131-generic root=/dev/mapper/ubunty00--vg-root ro recovery nomodeset

原始：

将末尾的ro recovery nomodeset修改为rw init=/bin/bash

按下 ctrl + x 组合键或者 F10 重新启动，进入单用户模式

3.5 直接修改密码

passwd root

两次输入密码，出现 password updated successfully ，则修改成功

按下 ctrl + alt + delete 组合键，直接重启系统，使用修改以后的密码正常登陆

到此已经修改root密码完成，赶快使用你修改的密码登录吧。
