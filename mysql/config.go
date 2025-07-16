// Package mysql @Author:冯铁城 [17615007230@163.com] 2025-07-15 11:53:10
package mysql

import "github.com/jmoiron/sqlx"

// Db mysql数据库连接对象
var Db *sqlx.DB

// DriverName mysql 数据库驱动名称
var DriverName = "mysql"

// Dsn mysql 数据库连接串
var Dsn = "root:root@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
