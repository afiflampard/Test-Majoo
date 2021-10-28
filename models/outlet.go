package models

import "gorm.io/gorm"

type Outlet struct {
	gorm.Model
	UserId uint   `gorm:"column:user_id" json:"user_id"`
	User   User   `gorm:"foreignKey:UserId"`
	Photo  string `gorm:"column:photo; type:varchar(255)" json:"photo"`
}
