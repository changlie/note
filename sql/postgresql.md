环境：ubuntu17
>  reference: http://www.ruanyifeng.com/blog/2013/12/getting_started_with_postgresql

## 安装
1. 首先，安装PostgreSQL客户端。
```
    sudo apt-get install postgresql-client
```
2.  然后，安装PostgreSQL服务器。
```
    sudo apt-get install postgresql
```
3. 想安装图形管理界面，可以运行下面命令
```
    sudo apt-get install pgadmin3
```

> PostgreSQL服务器会自动在本机的5432端口开启。

- 初次安装后，默认生成一个名为postgres的数据库和一个名为postgres的数据库用户。这里需要注意的是，同时还生成了一个名为postgres的Linux系统用户。

## 登录数据库
1. 切换到postgres用户。
```
    sudo su - postgres
```
2. 使用psql命令登录PostgreSQL控制台。
```
    psql
```
3. 第一件事是使用\password命令，为postgres用户设置一个密码。
```
    \password postgres
```

## 添加新用户
1. 新建一个 **Linux新用户** ，这里为`dbuser`。
```
    sudo adduser dbuser
```
2. 创建 **数据库用户**(该数据库用户名与linux用户名一致) dbuser（刚才创建的是Linux系统用户），并设置密码。
```
    CREATE USER dbuser WITH PASSWORD 'password';
``` 
   设置dbuser用户的密码，完成后退出控制台。
```
    \password dbuser
```
3. 创建用户数据库(该数据库名与用户名一致)，这里为dbuser，并指定所有者为dbuser。
```
    CREATE DATABASE exampledb OWNER dbuser;
```
4. 将exampledb数据库的所有权限都赋予dbuser，否则dbuser只能登录控制台，没有任何数据库操作权限。
```
    GRANT ALL PRIVILEGES ON DATABASE exampledb to dbuser;
``` 
5. 使用\q命令退出控制台（也可以直接按ctrl+D）。
```
    \q
```
6. 切换到dbuser, 然后用psql 重新进入数据库
```
   su dbuser
   psql
```

## 常用命令
- \h：查看SQL命令的解释，比如\h select。
- \?：查看psql命令列表。
- \l：列出所有数据库。
- \c [database_name]：连接其他数据库。
- \d：列出当前数据库的所有表格。
- \d [table_name]：列出某一张表格的结构。
- \du：列出所有用户。
- \e：打开文本编辑器。
- \conninfo：列出当前数据库和连接的信息。


