package models

import "time"

type User struct {
	ID        uint      `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id"`
	Username  string    `gorm:"column:username; type:varchar(255); not null; unique" json:"username"`
	Password  string    `gorm:"column:password; type:varchar(255); not null "json:"password"`
	FullName  string    `gorm:"column:fullname; type:varchar(255); not null "json:"fullname"`
	Photo     string    `gorm:"column:photo; type:varchar(255)"json:"photo"`
	CreatedAt time.Time `gorm:"column:created_at"json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at"json:"update_at"`
}
