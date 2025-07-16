// @Author:冯铁城 [17615007230@163.com] 2025-07-15 11:29:18
package main

import (
	"sqlx-demo/mysql"
	"sqlx-demo/sqlite"
)

// main函数 程序入口
func main() {

	//1.sqlite测试
	//sqliteTest()

	//2.mysql测试
	mysqlTest()
}

// sqlite测试
func sqliteTest() {

	//1.初始化sqlite数据库
	initSqlLiteDb(sqlite.DriverName, sqlite.DbFile)

	//2.确保最终数据库链接被关闭
	defer sqlite.Db.Close()

	//3.创建表
	sqlite.CreateTable()

	//4.插入数据
	sqlite.Insert()

	//5.查询数据
	sqlite.Select()

	//6.更新数据
	sqlite.Update()

	//7.删除数据
	sqlite.Delete()
}

func mysqlTest() {

	//1.初始化mysql数据库
	initMysqlDB(mysql.DriverName, mysql.Dsn)

	//2.创建表
	mysql.CreateTable()

	//4.插入数据
	mysql.Insert()

	//5.查询数据
	mysql.Select()

	//6.更新数据
	mysql.Update()

	////7.删除数据
	//mysql.Delete()
}
