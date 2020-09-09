### 1. 安装oracle即时客户端
1）下载地址： https://www.oracle.com/database/technologies/instant-client/linux-x86-64-downloads.html  
2）安装教程：
- https://oracle.github.io/odpi/doc/installation.html#oracle-instant-client-zip  
- https://godror.github.io/godror/doc/installation.html


### 2. go get github.com/godror/godror


### 3. demo
```golang
import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func Test_oracle(tt *testing.T) {
	// user/password@host:port/sid
    connStr := "system/oracle@192.168.1.103:1521/xe"
	db, err := sql.Open("godror", connStr)
	fmt.Println("open err:", err)
	defer db.Close()

	exec, err := db.Exec("insert into users(id, name) values(911, 'FBI')")
	fmt.Println("insert err1:", err)
	id, err := exec.LastInsertId()
	fmt.Println("insert result:", id, err)

	rows, err := db.Query("select id,name from users")
	fmt.Println("query err:", err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
```