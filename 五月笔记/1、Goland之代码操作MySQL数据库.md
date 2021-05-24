## Goland之代码操作MySQL数据库

### 一、安装MySQL驱动

```go
地址：https://github.com/Go-SQL-Driver/MySQL

import _ "github.com/go-sql-driver/mysql"
或者
go get github.com/go-sql-driver/mysql

-----基本连接-------
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 字符串的格式由对应的驱动进行定义
	// "username:password@tcp(IP:port)/database?charset=utf8mb4&loc=PRC&parseTime=true"
	dsn := "root:123456@tcp(192.168.241.66:3306)/mysql?charset=utf8mb4&loc=PRC&parseTime=true"
	a, err := sql.Open("mysql", dsn)
	fmt.Println(a, err)
}
```



### 二、代码操作

#### 2.1、基本操作

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 字符串的格式由对应的驱动进行定义
	// "username:password@tcp(IP:port)/database?charset=utf8mb4&loc=PRC&parseTime=true"
	dsn := "root:123456@tcp(192.168.241.66:3306)/goland?charset=utf8mb4&loc=PRC&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接失败，错误信息是：", err)
	}

	// 测试数据库连接状态
	fmt.Println("数据库连接是否成功", db.Ping())

	// 执行
	// 创建表
	resultt, _ := db.Exec(`
		create table if not exists testkk(
			id bigint primary key auto_increment,
			name varchar(32) not null default '' comment 'testkk name'
		) engine=innodb default charset utf8mb4;
	`)
	fmt.Println("执行成功", resultt)

	// 插入
	result, _ := db.Exec(`insert into testkk values(8,"kk")`)
	// 返回最后插入的数据
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// 更新、删除
	// db.Exec(`updata testkk set id = 2 `)

	// 查询
	rows, err := db.Query("select id, name from testkk")
	var (
		id   int
		name string
	)
	// 打印查询的数据，输出到id，name
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
```

