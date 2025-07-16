// Package mysql @Author:冯铁城 [17615007230@163.com] 2025-07-15 14:03:35
package mysql

import (
	"log"
)

// 删除表SQL
var dropTableSql = []string{
	"DROP TABLE IF EXISTS student_teacher",
	"DROP TABLE IF EXISTS students",
	"DROP TABLE IF EXISTS teachers",
}

// 建表SQL
var createTableSql = []string{
	"CREATE TABLE IF NOT EXISTS student_teacher (\n" +
		"    id INT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    student_id INT NOT NULL,\n" +
		"    teacher_id INT NOT NULL,\n" +
		"    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,\n" +
		"    UNIQUE KEY unique_student_teacher (student_id, teacher_id)\n" +
		");",
	"CREATE TABLE IF NOT EXISTS students (\n" +
		"    id INT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    name VARCHAR(100) NOT NULL,\n" +
		"    gender ENUM('male', 'female') DEFAULT 'male',\n" +
		"    age INT,\n" +
		"    email VARCHAR(255) UNIQUE,\n" +
		"    created_at DATETIME DEFAULT CURRENT_TIMESTAMP\n" +
		");",
	"CREATE TABLE IF NOT EXISTS teachers (\n" +
		"    id INT AUTO_INCREMENT PRIMARY KEY,\n" +
		"    name VARCHAR(100) NOT NULL,\n" +
		"    subject VARCHAR(100) NOT NULL,\n" +
		"    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,\n" +
		"    is_active TINYINT DEFAULT 1\n" +
		");",
}

// CreateTable 创建表
func CreateTable() {

	//1.执行删表SQL
	for _, sql := range dropTableSql {
		_, err := Db.Exec(sql)
		if err != nil {
			log.Fatal(err)
		}
	}

	//2.执行建表SQL
	for _, sql := range createTableSql {
		_, err := Db.Exec(sql)
		if err != nil {
			panic(err)
		}
	}

	//3.打印执行结果
	log.Printf("table create success")
}
