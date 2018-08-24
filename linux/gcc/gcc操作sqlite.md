> 参考: http://www.cnblogs.com/invisible2/p/9039192.html  
https://blog.csdn.net/zouleideboke/article/details/73649886  
http://www.cnblogs.com/hnrainll/archive/2011/09/08/2170490.html  
http://heyixian.iteye.com/blog/1998908

```shell
gcc -o demo main.c -lsqlite
```
```shell
./demo
```

```c
#include <sqlite3.h>
#include <stdlib.h>
#include <stdio.h>

void createTable();

int main(int argc, char **argv){
    sqlite3_os_init();
    char* filename = "hyx.db";
    sqlite3 *db = NULL;
    int rc;
    rc = sqlite3_open(filename, &db);
    if(rc){
    
    }else{
        printf("open sqlite file %s succeed; ", filename);
    }

    // create table users;
    //createTable(db);
    char *eMsg = 0;
    
    // insert
    char *sql1 = "insert into users(name,age) values('jay', 33)" ;
    sqlite3_exec(db, sql1, 0, 0, &eMsg);

    int nrow = 0, ncolumn = 0;
    char **azResult; //二维数组存放结果
    // query
    char *sql2 = "select * from users;";
    sqlite3_get_table(db, sql2, &azResult, &nrow, &ncolumn, &eMsg);
    printf("\n row:%d column=%d \n", nrow, ncolumn);
    int i=0;
    printf("\n The result of querying is: \n");
    for(i=0; i<(nrow+1)*ncolumn; i++)
        printf("azResult[%d] = %s \n", i, azResult[i]);
    
    sqlite3_close(db);
    return 1;
}

void createTable(sqlite3 *db){
    char *eMsg = 0;
    char *sql = "create table users( id integer primary key autoincrement, name varchar(55), age integer);";
    sqlite3_exec(db, sql, 0, 0, &eMsg);
}

```
