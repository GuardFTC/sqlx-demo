// Package model @Author:冯铁城 [17615007230@163.com] 2025-07-15 15:49:32
package model

type Teacher struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Subject   string `json:"subject" db:"subject"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"` // 默认值为当前日期
	IsActive  int    `json:"is_active,omitempty" db:"is_active"`   // 默认值为 1
}
