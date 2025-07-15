// Package sqlite @Author:冯铁城 [17615007230@163.com] 2025-07-15 15:47:43
package sqlite

import (
	"sqlx-demo/model"
)

// Insert 插入数据
func Insert() {

	//1.开启事务
	tx := Db.MustBegin()

	/*-----------------------------------------单个插入---------------------------------------------------*/
	//2.在事务内执行InsertSQL
	tx.Exec("INSERT INTO students(name,email,gender,age) VALUES(?,?,?,?)", "夏洛", "xialuo@xhs.com", "male", 18)
	tx.Exec("INSERT INTO students(name,email,gender,age) VALUES(?,?,?,?)", "马冬梅", "madongmei@xhs.com", "female", 17)

	//3.在事务内执行InsertSQL，基于结构体进行字段赋值
	teacher := new(model.Teacher)
	teacher.Name = "王老师"
	teacher.Subject = "语文"
	tx.NamedExec("INSERT INTO teachers(name,subject) VALUES(:name,:subject)", teacher)

	/*-----------------------------------------批量插入---------------------------------------------------*/
	//4.创建student切片
	students := []*model.Student{
		{Name: "袁华", Age: 18, Email: "yuanhua@xhs.com", Gender: "male"},
		{Name: "秋雅", Age: 18, Email: "qiuya@xhs.com", Gender: "female"},
		{Name: "张扬", Age: 18, Email: "zhangyang@xhs.com", Gender: "male"},
		{Name: "孟特", Age: 18, Email: "mengtejiao@xhs.com", Gender: "female"},
	}

	//5.批量插入
	tx.NamedExec("INSERT INTO students(name,email,gender,age) VALUES (:name,:email,:gender,:age)", students)

	//自定义版本批量插入，更加原始，但是适配性更强
	/*//5.定义初始SQL
	insertSql := "INSERT INTO students(name,email,gender,age) VALUES"

	//6.定义初始value切片
	var values []any

	//7.循环切片，拼接批量插入SQL
	for i, student := range students {

		//8.拼接SQL
		if i == len(students)-1 {
			insertSql += "(?,?,?,?)"
		} else {
			insertSql += "(?,?,?,?),"
		}

		//9.拼接value
		values = append(values, student.Name, student.Email, student.Gender, student.Age)
	}

	//10.执行批量插入SQL
	tx.Exec(insertSql, values...)*/

	//11.提交事务
	tx.Commit()
}
