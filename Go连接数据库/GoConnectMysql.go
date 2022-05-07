package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/users")
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
}
