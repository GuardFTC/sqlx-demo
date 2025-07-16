// @Author:冯铁城 [17615007230@163.com] 2025-07-15 11:38:36
package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"sqlx-demo/mysql"
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
	log.Printf("sqlite db init success\n")
}

// initMysqlDB 初始化Mysql数据库
func initMysqlDB(driverName string, dsn string) {

	//1.链接mysql数据库
	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	//2.赋值DB
	mysql.Db = db

	//3.日志打印
	log.Printf("mysql db init success\n")
}
