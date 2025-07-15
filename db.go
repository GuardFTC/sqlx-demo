// @Author:冯铁城 [17615007230@163.com] 2025-07-15 11:38:36
package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"sqlx-demo/sqlite"
)

// initSqlLiteDb 初始化Sqlite数据库
func initSqlLiteDb(driverName string, dbFile string) {

	//1.连接 SQLite 数据库文件（如不存在会自动创建）
	db, err := sqlx.Open(driverName, dbFile)
	if err != nil {
		log.Fatalln(err)
	}

	//2.赋值DB
	sqlite.Db = db

	//3.日志打印
	log.Printf("db init success\n")
}
