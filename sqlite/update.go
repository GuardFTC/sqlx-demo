// Package sqlite @Author:冯铁城 [17615007230@163.com] 2025-07-15 19:04:09
package sqlite

import (
	"log"

	"github.com/jmoiron/sqlx"
	"sqlx-demo/model"
)

// Update 更新数据
func Update() {

	/*-----------------------------------------根据主键ID更新---------------------------------------------------*/
	//1.查询数据
	var student model.Student
	Db.Get(&student, "select * from students where name = ?", "夏洛")

	//2.修改年龄
	student.Age = 30

	//3.更新数据
	tx := Db.MustBegin()
	exec, err := tx.NamedExec("update students set age = :age where name = :name", student)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	//4.打印受影响行数
	log.Println("更新数据后查询：----------------------------------------------------------------------------------------")
	affected, _ := exec.RowsAffected()
	log.Printf("更新数据成功，受影响的行数：%d\n", affected)

	//5.再次查询
	Db.Get(&student, "select * from students where name = ?", "夏洛")

	//6.打印信息
	log.Printf("id:%d,name:%s,age:%d,gender:%s,email:%s,created_at:%s\n",
		student.ID, student.Name, student.Age, student.Gender, student.Email, student.CreatedAt,
	)

	/*-----------------------------------------根据条件更新---------------------------------------------------*/
	//1.构建IN条件
	query, args, err := sqlx.In("update students set name = name || '_update' where age in (?) and id >= ?", []int{17, 18, 19}, 1)
	if err != nil {
		log.Fatal(err)
	}

	//2.rebind
	query = Db.Rebind(query)

	//3.根据条件批量更新
	tx = Db.MustBegin()
	exec, err = tx.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	log.Println("更新数据后查询：----------------------------------------------------------------------------------------")

	//5.构建IN条件
	query, args, err = sqlx.In("select * from students where age in (?) and id >= ?", []int{17, 18, 19}, 1)
	if err != nil {
		log.Fatal(err)
	}

	//6.rebind
	query = Db.Rebind(query)

	//7.根据条件查询
	var students []model.Student
	Db.Select(&students, query, args...)

	//8.打印受影响行数
	affected, _ = exec.RowsAffected()
	log.Printf("更新数据成功，受影响的行数：%d\n", affected)

	//10.遍历切片打印数据
	for _, student := range students {
		log.Printf("id:%d,name:%s,age:%d,gender:%s,email:%s,created_at:%s\n", student.ID, student.Name, student.Age, student.Gender, student.Email, student.CreatedAt)
	}
}
