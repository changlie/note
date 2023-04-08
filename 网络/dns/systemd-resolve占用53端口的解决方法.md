在Linux系统中有些软件（如：Dnsmasq解锁Netflix中的Dns等服务）可能要用到53的端口，但有些系统提示已使用（required port 53 already in use ）。
使用“netstat -tlunp|grep 53”查看53端口是不是systemd-resolved占用了（Oracle Cloud 永久免费 VPS/甲骨文免费VPS 就是这样）。如果是，下面为解决方法：

1、先停用 systemd-resolved 服务
```bash
systemctl stop systemd-resolved
```
2、编辑 /etc/systemd/resolved.conf 文件
```bash
vi /etc/systemd/resolved.conf
```
3、换下面说明更改，然后按一下“esc”键，再输入“:wq”（不要输入引号），回车保存即可。
```
[Resolve]
DNS=8.8.8.8  #取消注释，增加dns
#FallbackDNS=
#Domains=
#LLMNR=no
#MulticastDNS=no
#DNSSEC=no
#Cache=yes
DNSStubListener=no  #取消注释，把yes改为no
```
4、最后运行下面命令即可。
```bash
ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
```