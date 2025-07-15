// Package sqlite @Author:冯铁城 [17615007230@163.com] 2025-07-15 16:37:22
package sqlite

import (
	"log"

	"sqlx-demo/model"
)

// Select 查询数据
func Select() {

	/*-----------------------------------------批量查询数据-----------------------------------------------*/
	//1.创建结构体切片
	var students []model.Student

	//2.执行查询
	err := Db.Select(&students, "select * from students where age = $1 order by id desc", 18)
	if err != nil {
		log.Fatal(err)
	}

	//3.遍历结构体切片
	log.Println("批量查询数据：----------------------------------------------------------------------------------------")
	for _, student := range students {
		log.Printf("id:%d, name:%s, age:%d, gender:%s, email:%s, created_at:%s\n", student.ID, student.Name, student.Age, student.Gender, student.Email, student.CreatedAt)
	}

	/*-----------------------------------------单个查询数据-----------------------------------------------*/
	//1.创建结构体
	var teacher model.Teacher

	//2.执行查询
	err = Db.Get(&teacher, "select * from teachers where id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("单个查询数据：----------------------------------------------------------------------------------------")
	log.Printf("id:%d,name:%s,subject:%s,is_active:%d,created_at:%s\n",
		teacher.ID, teacher.Name, teacher.Subject, teacher.IsActive, teacher.CreatedAt,
	)

	/*-----------------------------------------批量查询数据，使用单个结构体接收--------------------------------*/
	//1.创建结构体
	var student model.Student

	//2.执行查询
	rows, err := Db.Queryx("select * from students where age = $1 order by id desc", 18)
	if err != nil {
		log.Fatal(err)
	}

	//3.遍历rows
	log.Println("批量查询数据：----------------------------------------------------------------------------------------")
	students = []model.Student{}
	for rows.Next() {

		//4.将查询结果赋给student
		err := rows.StructScan(&student)
		if err != nil {
			log.Fatal(err)
		}

		//5.存入切片
		students = append(students, student)
	}

	//6.遍历切片打印
	for _, student = range students {
		log.Printf("id:%d,name:%s,age:%d,gender:%s,email:%s,created_at:%s\n",
			student.ID, student.Name, student.Age, student.Gender, student.Email, student.CreatedAt,
		)
	}

	/*-----------------------------------------批量查询数据，通过结构体传参--------------------------------*/
	//1.创建参数结构体
	paramStudent := model.Student{}
	paramStudent.Age = 17

	//2.执行查询
	rows, err = Db.NamedQuery("select * from students where age = :age order by id desc", paramStudent)
	if err != nil {
		log.Fatal(err)
	}

	//3.遍历rows
	log.Println("批量查询数据：----------------------------------------------------------------------------------------")
	for rows.Next() {

		//4.将查询结果赋给student
		err = rows.StructScan(&student)
		if err != nil {
			log.Fatal(err)
		}

		//5.打印结果
		log.Printf("id:%d,name:%s,age:%d,gender:%s,email:%s,created_at:%s\n",
			student.ID, student.Name, student.Age, student.Gender, student.Email, student.CreatedAt,
		)
	}
}
