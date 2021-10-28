package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TanggalBeli time.Time  `gorm:"column:tanggal_beli" json:"tanggal_beli"`
	IDPembeli   uint       `gorm:"column:id_pembeli" json:"id_pembeli"`
	IDOutlet    uint       `gorm:"column:id_outlet" json:"id_outlet"`
	Outlet      Outlet     `gorm:"foreignKey:IDOutlet"`
	Pembeli     User       `gorm:"foreignKey:IDPembeli"`
	IDProduct   uint       `gorm:"column:id_product" json:"id_product"`
	Product     Product    `gorm:"foreignKey:IDProduct"`
	Total       uint       `gorm:"column:total" json:"total"`
	NoState     uint       `gorm:"column:no_state" json:"no_state"`
	OrderState  OrderState `gorm:"foreignKey:NoState"`
}

type OrderState struct {
	gorm.Model
	State string `gorm:"column:state" json:"state"`
}

type History struct {
	gorm.Model
	IDOrder    uint       `gorm:"column:order_id" json:"order_id"`
	Order      Order      `gorm:"foreignKey:IDOrder"`
	IDOutlet   uint       `gorm:"column:outlet_id" json:"outlet_id"`
	Outlet     Outlet     `gorm:"foreignKey:IDOutlet"`
	NoState    uint       `gorm:"column:stateNo" json:"no_state"`
	OrderState OrderState `gorm:"foreignKey:NoState"`
}
