package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	NamaProduct string `gorm:"column:nama_product" json:"nama_product"`

	IdUser       uint   `gorm:"column: id_user" json:"id_user"`
	User         User   `gorm:"foreignKey:IdUser"`
	IdOutlet     uint   `gorm:"column:id_outlet" json:"idOutlet"`
	Outlet       Outlet `gorm:"foreignKey:IdOutlet"`
	Foto         string `gorm:"column:photo" json:"photo"`
	HargaProduct uint   `gorm:"column:harga" json:"harga"`
	MaxStock     uint   `gorm:"column:max_stock" json:"max_stock"`
	Stock        uint   `gorm:"column:stock" json:"stock"`
}
