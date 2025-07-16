// Package sqlite @Author:冯铁城 [17615007230@163.com] 2025-07-15 20:11:02
package sqlite

import (
	"log"

	"github.com/jmoiron/sqlx"
	"sqlx-demo/model"
)

func Delete() {

	/*-----------------------------------------根据主键ID删除---------------------------------------------------*/
	//1.创建结构体
	var teacher model.Teacher

	//2.根据主键ID查询
	err := Db.Get(&teacher, "select * from teachers where id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}

	//3.删除
	tx := Db.MustBegin()
	exec, err := tx.NamedExec("delete from teachers where id = :id", teacher)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	//4.打印受影响行数
	affected, _ := exec.RowsAffected()
	log.Printf("删除数据成功，受影响的行数：%d\n", affected)

	//5.再次查询，因为已经被删除了，所以查询不到数据
	err = Db.Get(&teacher, "select * from teachers where id = $1", 1)
	if err != nil {
		log.Println(err)
	}

	/*-----------------------------------------根据主键ID批量删除-------------------------------------------------*/
	//1.构建IN查询条件
	query, args, err := sqlx.In("delete from students where id in (?)", []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}

	//2.rebind
	query = Db.Rebind(query)

	//3.删除
	tx = Db.MustBegin()
	exec, err = tx.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	//4.打印受影响行数
	affected, _ = exec.RowsAffected()
	log.Printf("删除数据成功，受影响的行数：%d\n", affected)

	//5.构建IN查询条件
	query, args, err = sqlx.In("select * from students where id in (?)", []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}

	//6.rebind
	query = Db.Rebind(query)

	//7.查询
	var students []model.Student
	err = Db.Select(&students, query, args...)
	if err != nil {
		log.Println(err)
	}
	if students == nil {
		log.Println("经过验证，数据删除成功")
	}

	/*-----------------------------------------根据条件批量删除-------------------------------------------------*/
	//1.删除
	tx = Db.MustBegin()
	exec, err = tx.Exec("delete from students where gender = $1", "female")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	//2.打印受影响行数
	affected, _ = exec.RowsAffected()
	log.Printf("删除数据成功，受影响的行数：%d\n", affected)

	//3.查询验证
	err = Db.Select(&students, "select * from students where gender = $1", "female")
	if err != nil {
		log.Println(err)
	}
	if students == nil {
		log.Println("经过验证，数据删除成功")
	}
}
