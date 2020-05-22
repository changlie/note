### 字符串占位符
```golang
res := os.Expand("hi! ${name}, age:${age}, come from $addr.  $unknow.", mapping)
fmt.Println(res)
func mapping(key string) string {
  fmt.Println("call by other func")
  m := map[string]string{
    "name":"jerry",
    "age":"33",
    "addr":"china",
  }
  if key != "" {
    val, ok := m[key]
    if ok {
      return val
    }
    return key
  }

  return key
}
```
输出
```
hi! jerry, age:33, come from china.  unknow.
```

### 环境变量相关（所有操作都不会持久化）
```golang
func Clearenv()   // 清除环境变量（只限这次启动的会话，相关于缓存式清空）
func Environ() []string // 获取所有环境变量
func LookupEnv(key string) (string, bool)  // 获取单个环境变量
func Getenv(key string) string // 获取单个环境变量
func Setenv(key, value string) error // 设置环境变量
func Unsetenv(key string) error // 删除单个环境变量
```
### 文件操作
```golang
// 创建目录
func Mkdir(name string, perm FileMode) error  // 一次创建一个新目录
func MkdirAll(path string, perm FileMode) error // 同时创建多个新目录
// 删除文件或目录
func Remove(name string) error  // removes the named file or (empty) directory.
func RemoveAll(path string) error // removes path and any children it contains.
// 重命名或移动文件
func Rename(oldpath, newpath string) error // renames (moves) oldpath to newpath. 如果newpath是一个已存在的文件，则进行覆盖操作
// 判断是否为同一个文件
func SameFile(fi1, fi2 FileInfo) bool 

// 文件读写操作
type File
    func Create(name string) (*File, error)
    func NewFile(fd uintptr, name string) *File
    func Open(name string) (*File, error)  // 读文件
    func OpenFile(name string, flag int, perm FileMode) (*File, error)  // 允许追加写文件
    func (f *File) Chdir() error
    func (f *File) Chmod(mode FileMode) error  // 更改权限
    func (f *File) Chown(uid, gid int) error // changes the numeric uid and gid of the named file
    func (f *File) Close() error  // closes the File, rendering it unusable for I/O(关闭文件，使无法再对其进行I/O操作)
    func (f *File) Fd() uintptr
    func (f *File) Name() string  // 获取文件名
    func (f *File) Read(b []byte) (n int, err error) // reads up to len(b) bytes from the File
    func (f *File) ReadAt(b []byte, off int64) (n int, err error) // reads len(b) bytes from the File starting at byte offset off
    func (f *File) Readdir(n int) ([]FileInfo, error)  // 获取目录下的文件信息
    func (f *File) Readdirnames(n int) (names []string, err error) // 获取目录下的文件信息
    func (f *File) Seek(offset int64, whence int) (ret int64, err error) // The behavior of Seek on a file opened with O_APPEND is not specified.
    func (f *File) SetDeadline(t time.Time) error
    func (f *File) SetReadDeadline(t time.Time) error
    func (f *File) SetWriteDeadline(t time.Time) error
    func (f *File) Stat() (FileInfo, error)  // 文件名，大小，判断是否为目录，
    func (f *File) Sync() error // commits the current contents of the file to stable storage.相当于java的flush()
    func (f *File) SyscallConn() (syscall.RawConn, error)
    func (f *File) Truncate(size int64) error // changes the size of the file. It does not change the I/O offset.
    func (f *File) Write(b []byte) (n int, err error)  writes len(b) bytes to the File
    func (f *File) WriteAt(b []byte, off int64) (n int, err error) // writes len(b) bytes to the File starting at byte offset off
    func (f *File) WriteString(s string) (n int, err error)
type FileInfo
    func Lstat(name string) (FileInfo, error)
    func Stat(name string) (FileInfo, error)
type FileMode
    func (m FileMode) IsDir() bool
    func (m FileMode) IsRegular() bool
    func (m FileMode) Perm() FileMode
    func (m FileMode) String() string
```

type FileInfo    
A FileInfo describes a file and is returned by Stat and Lstat.
```golang
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
```


>  参考：

[os包 腾讯云社区](https://cloud.tencent.com/developer/article/1342799)

[官方文档](https://golang.org/pkg/os/)
