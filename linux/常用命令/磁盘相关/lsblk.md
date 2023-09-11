列出块设备的信息。
```bash
$ lsblk
NAME        MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT
loop0         7:0    0     4K  1 loop /snap/bare/5
loop1         7:1    0 118.2M  1 loop /snap/core/15511
loop2         7:2    0 118.2M  1 loop /snap/core/15419
loop3         7:3    0  55.7M  1 loop /snap/core18/2785
loop4         7:4    0  55.7M  1 loop /snap/core18/2790
loop5         7:5    0  73.9M  1 loop /snap/core22/766
loop6         7:6    0  63.5M  1 loop /snap/core20/2015
loop7         7:7    0 466.6M  1 loop /snap/gnome-42-2204/111
loop8         7:8    0  73.9M  1 loop /snap/core22/858
loop9         7:9    0 139.4M  1 loop /snap/docker/2746
loop10        7:10   0  63.5M  1 loop /snap/core20/1950
loop11        7:11   0 485.5M  1 loop /snap/gnome-42-2204/126
loop12        7:12   0 128.9M  1 loop /snap/docker/2893
loop13        7:13   0 349.7M  1 loop /snap/gnome-3-38-2004/140
loop14        7:14   0 349.7M  1 loop /snap/gnome-3-38-2004/143
loop15        7:15   0  91.7M  1 loop /snap/gtk-common-themes/1535
loop16        7:16   0    46M  1 loop /snap/snap-store/638
loop17        7:17   0  53.3M  1 loop /snap/snapd/19457
loop18        7:18   0  12.3M  1 loop /snap/snap-store/959
loop19        7:19   0  40.9M  1 loop /snap/snapd/19993
sda           8:0    0 931.5G  0 disk 
├─sda1        8:1    0   100M  0 part /boot/efi
├─sda2        8:2    0    16M  0 part 
├─sda3        8:3    0 315.8G  0 part 
└─sda4        8:4    0 615.6G  0 part /
nvme0n1     259:0    0   477G  0 disk /media/d
├─nvme0n1p1 259:3    0   255G  0 part 
└─nvme0n1p2 259:4    0   210G  0 part 
````

### 创建文件格式
sudo mkfs -t ext4 /dev/nvme0n1p1