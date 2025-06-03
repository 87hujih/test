package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入但不直接使用，用于注册驱动
)

func main() {
	// 数据库连接信息
	dbUser := "root"
	dbPass := "654321"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "one"

	// 构建连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // 无法连接数据库
	}
	defer db.Close() // 确保在函数结束时关闭数据库连接

	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // 无法 ping 通数据库
	}

	fmt.Println("成功连接到 MySQL 数据库")

	// 在这里进行 CRUD 操作
}
