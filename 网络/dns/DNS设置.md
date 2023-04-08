1. 修改`/etc/systemd/resolved.conf` 文件
修改 `DNS`及`DNSStubListener`的值
```
[Resolve]
DNS=192.168.1.105
#FallbackDNS=
#Domains=
#LLMNR=no
#MulticastDNS=no
#DNSSEC=no
#DNSOverTLS=no
#Cache=no-negative
DNSStubListener=no
#ReadEtcHosts=yes
```
2. 重启dns服务
```
sudo systemctl restart systemd-resolved
sudo systemctl enable systemd-resolved
```
3. 更新dns地址
```
sudo mv /etc/resolv.conf /etc/resolv.conf.bak
sudo ln -sf /run/systemd/resolve/resolv.conf /etc/
```
3. 查看当前dns信息
```
systemd-resolve --status
或
resolvectl status
```


### 最终脚本
```bash
sed -e 's/\#\(\DNS=\)/\1192\.168\.1\.105/g' -e 's/\#\(\DNSStubListener=\)yes/\1no/' /etc/systemd/resolved.conf > .ttmp
sudo cat .ttmp > /etc/systemd/resolved.conf
rm .ttmp

sudo systemctl restart systemd-resolved
sudo systemctl enable systemd-resolved

sudo mv /etc/resolv.conf /etc/resolv.conf.bak
sudo ln -sf /run/systemd/resolve/resolv.conf /etc/

systemd-resolve --status
```




