// @Author:冯铁城 [17615007230@163.com] 2025-07-15 11:38:36
package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"sqlx-demo/mysql"
	"sqlx-demo/sqlite"
)

// initSqlLiteDb 初始化Sqlite数据库
func initSqlLiteDb(driverName string, dbFile string) {

	//1.初始化数据库链接配置(只准备不连接),SQLite数据库文件,如不存在会自动创建
	db, err := sqlx.Open(driverName, dbFile)
	if err != nil {
		log.Fatalln(err)
	}

	//2.通过ping立即创建链接，验证数据库能否链接成功
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	//3.设置连接池参数
	db.SetMaxOpenConns(20)                 // 最多20个连接
	db.SetMaxIdleConns(10)                 // 最多10个空闲连接
	db.SetConnMaxLifetime(5 * time.Minute) // 每个连接最多用5分钟
	db.SetConnMaxIdleTime(1 * time.Minute) // 空闲超过1分钟就关闭

	//4.赋值DB
	sqlite.Db = db

	//5.日志打印
	log.Printf("sqlite db init success\n")
}

// initMysqlDB 初始化Mysql数据库
func initMysqlDB(driverName string, dsn string) {

	//1.链接mysql数据库
	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	//2.设置连接池参数
	db.SetMaxOpenConns(20)                 // 最多20个连接
	db.SetMaxIdleConns(10)                 // 最多10个空闲连接
	db.SetConnMaxLifetime(5 * time.Minute) // 每个连接最多用5分钟
	db.SetConnMaxIdleTime(1 * time.Minute) // 空闲超过1分钟就关闭

	//3.赋值DB
	mysql.Db = db

	//4.日志打印
	log.Printf("mysql db init success\n")
}
