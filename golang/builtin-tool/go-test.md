> 参考:  [golang test测试使用](https://studygolang.com/articles/2491)

1. 创建测试文件夹mysql，文件夹下的go文件的package必须与文件夹名一致（不然会识别不到）

2. 创建需要测试的文件mysql.go(使用github.com/go-sql-driver/mysql包)

```golang
package mysql

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func findByPk(pk int) int {
    var num int = 0
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/plugin_master?charset=utf8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    stmtOut, err := db.Prepare("select id from t_admin where id=?")
    if err != nil {
        panic(err.Error())
    }
    defer stmtOut.Close()

    err = stmtOut.QueryRow(pk).Scan(&num)
    if err != nil {
        panic(err.Error())
    }
    return num
}
```

3. 创建单元测试用例文件`mysql_test.go`(文件名必须是`*_test.go`的类型，``*``代表要测试的文件名，函数名必须以`Test`开头如：`TestXxx`或`Test_xxx`)

```golang
package mysql

import (
    "testing"
)

func Test_findByPk(t *testing.T) {
    num := findByPk(1)
    t.Log(num)
}
```

测试所有的文件 `go test`，将对当前目录下的所有``*_test.go``文件进行编译并自动运行测试。

- ``-file``参数: 测试某个文件使用. `go test –file *.go` 。   
例如：`go test -file mysql_test.go`，"-file"参数不是必须的，可以省略，如果你输入`go test b_test.go`也会得到一样的效果。

- ``-run`` 参数: 测试某个方法 `go test -run='Test_xxx'`.

- `-v` 参数: 显示所以用例测试结果.
`go test -v ...` 表示无论用例是否测试通过都会显示结果，不加"-v"表示只显示未通过的用例结果

4. 创建benchmark性能测试用例文件`mysql_b_test.go`(文件名必须是``*_b_test.go``的类型，``*``代表要测试的文件名，函数名必须以Benchmark开头如：BenchmarkXxx或Benchmark_xxx)

```golang
package mysql

import (
    "testing"
)

func Benchmark_findByPk(b *testing.B) {
    for i := 0; i < b.N; i++ { //use b.N for looping
        findByPk(1)
    }
}
```
进行所有go文件的benchmark测试 `go test -bench=".*"` 或 `go test . -bench=".*"`   
对某个go文件进行benchmark测试 `go test mysql_b_test.go -bench=".*"`

5. 用性能测试生成CPU状态图（暂未测试使用）

使用命令：
`go test -bench=".*" -cpuprofile=cpu.prof -c`
cpuprofile是表示生成的cpu profile文件  
-c是生成可执行的二进制文件，这个是生成状态图必须的，它会在本目录下生成可执行文件mysql.test
然后使用go tool pprof工具
```
go tool pprof mysql.test cpu.prof
```
调用web（需要安装graphviz）来生成svg文件，生成后使用浏览器查看svg文件   
参考 http://www.cnblogs.com/yjf512/archive/2013/01/18/2865915.html
