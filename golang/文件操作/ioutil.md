```golang
func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadDir(dirname string) ([]os.FileInfo, error)
func ReadFile(filename string) ([]byte, error)
func TempDir(dir, pattern string) (name string, err error)
func TempFile(dir, pattern string) (f *os.File, err error)
func WriteFile(filename string, data []byte, perm os.FileMode) error
```

### 常用方法
#### 1.读取文件（一次全部读取）
```golang
func ReadAll(r io.Reader) ([]byte, error)
func ReadFile(filename string) ([]byte, error)
```
ReadFile 比 ReadAll高效    
ReadAll 的输入不一定是文件，也可以字符串转Reader,如下：
```golang
s := strings.NewReader("Hello World!")
ra, _ := ioutil.ReadAll(s)
fmt.Printf("%s", ra)
// Hello World!
```

#### 2. 获取目录下的所有文件信息
```golang
func ReadDir(dirname string) ([]os.FileInfo, error)
```

#### 3. 向文件写东西
> 不能追加写
```golang
func WriteFile(filename string, data []byte, perm os.FileMode) error
```
