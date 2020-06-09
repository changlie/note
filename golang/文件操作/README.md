### 文件相关


获取当前工程根路径（打包时会固定死）
```golang
str, _ := os.Getwd() // 获取工作路径 (wd: work directory)
os.Chdir("d:/work") // 更改工作路径
```

获取当前路径（指向执行文件）
```golang
dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
或
dir, _ := filepath.Dir(os.Args[0])
```

按任意键继续 的功能实现 (允许程序在window下，执行完成后不会马上关闭)
```golang
func pause() {
	fmt.Println("press any key to continue...")
	var input string
	for {
		fmt.Scan(&input)
		fmt.Println("input: ", input)
	}
}
```
