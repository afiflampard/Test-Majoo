package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username; type:varchar(255); not null; unique" json:"username"`
	Password string `gorm:"column:password; type:varchar(255); not null" json:"password"`
	FullName string `gorm:"column:fullname; type:varchar(255); not null" json:"fullname"`
	RoleID   uint   `gorm:"role_id" json:"roleId"`
	Role     Role   `gorm:"foreignKey:RoleID"`
	Photo    string `gorm:"column:photo; type:varchar(255)" json:"photo"`
}

type Role struct {
	gorm.Model
	Role string `gorm:"column:role; type:varchar(255)" json:"role"`
}
