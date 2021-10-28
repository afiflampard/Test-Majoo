package service

import (
	"errors"
	"majoo/entities"
	"majoo/models"
	"time"

	"gorm.io/gorm"
)

func TransaksiJual(db *gorm.DB, order entities.OrderRequest) (*models.Order, error) {
	tx := db.Begin()
	var product models.Product
	if err := tx.Find(&product, order.IDProduct).Error; err != nil {
		return nil, err
	}
	if product.Stock-order.Jumlah < 0 {
		return nil, errors.New("Jummlah tidak memenuhi")
	}
	orderProduct := models.Order{
		TanggalBeli: time.Now(),
		IDPembeli:   order.IDPembeli,
		IDOutlet:    product.IdOutlet,
		IDProduct:   product.ID,
		Total:       (product.HargaProduct * order.Jumlah),
		NoState:     1,
	}
	if err := tx.Debug().Create(&orderProduct).Error; err != nil {

		return nil, err
	}
	product.Stock = product.Stock - order.Jumlah
	if err := db.Save(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	newProduct := models.Product{
		IdOutlet:     product.IdOutlet,
		NamaProduct:  product.NamaProduct,
		IdUser:       uint(order.IDPembeli),
		HargaProduct: product.HargaProduct,
		MaxStock:     order.Jumlah,
		Stock:        order.Jumlah,
	}
	if err := tx.Create(&newProduct).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &orderProduct, nil
}

func HistoryPenjualan(db *gorm.DB, idOutlet int) (*[]models.Order, error) {
	var history []models.Order
	if err := db.Where("id_outlet = ?  AND no_state = ?", idOutlet, 1).Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func HistoryPembelian(db *gorm.DB, idOutlet int) (*[]models.Order, error) {
	var history []models.Order
	if err := db.Where("id_outlet = ?  AND no_state = ?", idOutlet, 2).Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func TransaksiBeli(db *gorm.DB, order entities.OrderRequest) (*models.Order, error) {
	tx := db.Begin()
	var product models.Product
	if err := tx.Find(&product, order.IDProduct).Error; err != nil {
		return nil, err
	}
	if product.Stock-order.Jumlah < 0 {
		return nil, errors.New("Jummlah tidak memenuhi")
	}
	orderProduct := models.Order{
		TanggalBeli: time.Now(),
		IDPembeli:   order.IDPembeli,
		IDOutlet:    product.IdOutlet,
		IDProduct:   product.ID,
		Total:       (product.HargaProduct * order.Jumlah),
		NoState:     2,
	}
	if err := tx.Create(&orderProduct).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	product.Stock = product.Stock - order.Jumlah
	if err := db.Save(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	newProduct := models.Product{
		IdOutlet:     product.IdOutlet,
		NamaProduct:  product.NamaProduct,
		IdUser:       uint(order.IDPembeli),
		HargaProduct: product.HargaProduct,
		MaxStock:     order.Jumlah,
		Stock:        order.Jumlah,
	}
	if err := tx.Create(&newProduct).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &orderProduct, nil
}
