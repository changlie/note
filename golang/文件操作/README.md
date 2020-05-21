### 文件相关


获取当前工程根路径（打包时会固定死）
```golang
str, _ := os.Getwd()
```

获取当前路径（指向执行文件）
```golang
dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
```
