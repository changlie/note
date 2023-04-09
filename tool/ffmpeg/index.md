## 视频截取
ffmpeg -ss 00:00:00 -i a.mkv -to 01:10:00 -c:v copy -c:a copy  f.mkv
## 视频合并
ffmpeg -f concat -i filelist.txt -c copy output.mkv
`filelist.txt`
```
file 'a.mkv'
file 'b.mkv'
```
## mkv 转 mp4

ffmpeg -i a.mkv -vcodec copy -acodec copy a.mp4
