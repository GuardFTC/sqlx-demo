// Package sqlite @Author:冯铁城 [17615007230@163.com] 2025-07-15 14:03:35
package sqlite

import (
	"log"
)

// 删除表SQL
const dropTableSql = `
drop table student_teacher;
drop table students;
drop table teachers;
`

// 建表SQL
const createTableSql = `
create table IF NOT EXISTS student_teacher
(
    id         INTEGER
        primary key autoincrement,
    student_id INTEGER not null,
    teacher_id INTEGER not null,
	created_at TEXT default (datetime('now', '+8 hours')),
    unique (student_id, teacher_id)
);

create table IF NOT EXISTS students
(
    id         INTEGER
        primary key autoincrement,
    name       TEXT not null,
    gender     TEXT default 'male',
    age        INTEGER,
    email      TEXT
        unique,
	created_at TEXT default (datetime('now', '+8 hours')),
    check (gender IN ('male', 'female'))
);

create table IF NOT EXISTS teachers
(
    id         INTEGER
        primary key autoincrement,
    name       TEXT not null,
    subject    TEXT not null,
	created_at TEXT default (datetime('now', '+8 hours')),
    is_active  INTEGER default 1
);
`

// CreateTable 创建表
func CreateTable() {

	//1.执行删表SQL
	_, err := Db.Exec(dropTableSql)
	if err != nil {
		log.Fatal(err)
	}

	//2.执行建表SQL
	_, err = Db.Exec(createTableSql)
	if err != nil {
		panic(err)
	}

	//3.打印执行结果
	log.Printf("table create success")
}
