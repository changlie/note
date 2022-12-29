
## 宿主ubuntu20 安装  vmware16

```bash
# 第一步：基本资源安装
chmod +x VMware-Workstation-Full-16.2.1-18811642.x86_64.bundle
sudo ./VMware-Workstation-Full-16.2.1-18811642.x86_64.bundle

# 第二步：安装依赖
sudo apt install gcc
sudo apt install make
sudo apt-get install libaio1 libglib2.0-dev

# 第三步: 安装模块1 ，若全部成功则无需进行第四步
sudo vmware-modconfig --console --install-all

# 第四步: 安装模块2， 安装
git clone https://github.com/mkubecek/vmware-host-modules.git
#上面的GitHub地址有可能用不了，使用下面GitHub的镜像网站
cd vmware-host-modules
git checkout workstation-16.2.1  #VMware对应的版本号
make
make install
# 最后重启即可
```






