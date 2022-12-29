## 查看 cpu信息
```bash
lscpu
```

## 查看 cpu温度
1. 方式一
```bash
# sudo apt install lm-sensors (如果没有安装)
sensors
```
2. 方式二
```bash
watch -n 0.1 echo CPU: $[$(cat /sys/class/thermal/thermal_zone0/temp)/1000]°
```

## 查看 cpu 实时频率
```bash
watch grep \"cpu MHz\" /proc/cpuinfo
```