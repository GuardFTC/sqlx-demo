// Package sqlite @Author:冯铁城 [17615007230@163.com] 2025-07-15 11:53:10
package sqlite

import "github.com/jmoiron/sqlx"

// Db Sqlite数据库连接对象
var Db *sqlx.DB

// DriverName Sqlite 数据库驱动名称
var DriverName = "sqlite3"

// DbFile 数据库文件
var DbFile = "D:\\base\\sqllite\\sqlx.sqlite3"
