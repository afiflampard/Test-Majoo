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
		IDProduct:   product.ID,
		Total:       (product.HargaProduct * order.Jumlah),
		NoState:     1,
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
	history := []models.History{
		{
			IDOrder:   orderProduct.ID,
			IDpembeli: order.IDPembeli,
			IDOutlet:  product.IdOutlet,
			IdProduct: product.ID,
			NoState:   1,
		},
		{
			IDOrder:   orderProduct.ID,
			IDpembeli: order.IDPembeli,
			IDOutlet:  product.IdOutlet,
			IdProduct: product.ID,
			NoState:   2,
		},
	}
	for _, data := range history {
		if err := db.Create(&data).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()

	return &orderProduct, nil
}

func HistoryPenjualan(db *gorm.DB, idOutlet int) (*[]models.History, error) {
	var history []models.History
	if err := db.Where("id_outlet = ?  AND no_state = ?", idOutlet, 1).Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func HistoryPembelian(db *gorm.DB, idOutlet int) (*[]models.History, error) {
	var history []models.History
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
	history := []models.History{
		{
			IDOrder:   orderProduct.ID,
			IDpembeli: order.IDPembeli,
			IDOutlet:  product.IdOutlet,
			IdProduct: product.ID,
			NoState:   2,
		},
	}
	for _, data := range history {
		if err := db.Create(&data).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()

	return &orderProduct, nil
}
