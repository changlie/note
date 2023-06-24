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

