## 视频截取
ffmpeg -ss 00:00:00 -i a.mkv -to 01:10:00 -c:v copy -c:a copy  f.mkv

ffmpeg -ss 00:00:20 -i a.mkv -to 02:00:00 -c:v copy -c:a copy  f.mkv

## 视频合并
ffmpeg -f concat -i filelist -c copy output.mkv
`filelist`
```
file 'a.mkv'
file 'b.mkv'
```

## mkv 转 mp4

ffmpeg -i a.mkv -vcodec copy -acodec copy a.mp4

## 截取视频帧作为图片

1. 生产视频的每分钟的缩略图：

```bash
ffmpeg -i film.mp4 -vf fps=1/60 img%03d.jpg
```

上面这个-vf fps=1/60，就是1分钟的意思，比如说视频25分钟，会生成25个jpg图片，分别是img001，img002....，

当fps=1的时候：就代表每一秒截取个缩略图  
而fps=1/600：则代表每10分钟截取一个画面

2. 截取视频任意帧图片
```
ffmpeg -i ./xxx.mkv  -y -f image2 -vframes 1 -ss 430 ./preview.jpg
```

**参数说明**
```
参数	参数值
-i	视频路径
-y	覆盖输出文件
-f	强迫采用格式fmt
-ss	开始截图时间 seconds or in hh:mm:ss[.xxx] 如果截图开始时间越接近篇尾，所花费的时间就会越长
-vframes	截取帧
```