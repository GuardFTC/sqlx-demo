// Package model @Author:冯铁城 [17615007230@163.com] 2025-07-15 15:49:32
package model

// Student 学生结构体
type Student struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Gender    string `json:"gender,omitempty" db:"gender"` // 默认值为 "male"
	Age       int    `json:"age,omitempty" db:"age"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"` // 默认值为当前时间
}
