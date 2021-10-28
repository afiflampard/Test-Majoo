package service

import (
	"errors"
	"fmt"
	"io"
	"log"
	"majoo/entities"
	"majoo/models"
	"mime/multipart"
	"os"

	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product entities.RequestProduct, id int, idUser int) (*models.Product, error) {
	var productTemp models.Product
	fmt.Println("Id nya", id)
	var (
		user models.User
	)
	if err := db.Debug().First(&user, idUser).Preload("Role").Error; err != nil {
		return nil, errors.New("User Not Found")
	}
	var count int64
	db.Where("nama_product = ? AND id_outlet = ?", product.NamaProduk, id).Preload("Outlet").Preload("Outlet.User").Preload("Outlet.User.Role").Find(&productTemp).Count(&count)

	if user.RoleID == 1 {
		if count < 1 {
			addProduct := models.Product{
				IdOutlet:    uint(id),
				NamaProduct: product.NamaProduk,

				IdUser:       uint(idUser),
				HargaProduct: product.Harga,
				MaxStock:     product.Stock,
				Stock:        product.Stock,
			}
			if err := db.Create(&addProduct).Error; err != nil {
				return nil, err
			}
			return &addProduct, nil
		}
		productTemp.HargaProduct = product.Harga
		productTemp.NamaProduct = product.NamaProduk
		productTemp.Stock = product.Stock
		if err := db.Save(&productTemp).Error; err != nil {
			return nil, err
		}

	} else if user.RoleID == 2 {
		addProduct := models.Product{
			IdUser:       user.ID,
			IdOutlet:     1,
			NamaProduct:  product.NamaProduk,
			HargaProduct: product.Harga,
			MaxStock:     product.Stock,
			Stock:        product.Stock,
		}
		if err := db.Create(&addProduct).Error; err != nil {
			return nil, err
		}
		return &addProduct, nil
	}
	return &productTemp, nil
}

func GetProductByID(db *gorm.DB, idProduct int) (*models.Product, error) {
	var product models.Product
	if err := db.Debug().Where("id = ?", idProduct).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetAllProduct(db *gorm.DB, idOutlet int) ([]models.Product, error) {
	var products []models.Product
	if err := db.Where("id_outlet = ?", idOutlet).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}
func DeleteProduct(db *gorm.DB, idProduct int) (map[string]string, error) {
	var product models.Product
	if err := db.Where("id = ?", idProduct).Delete(&product).Error; err != nil {
		return map[string]string{
			"message": "Product tidak ada",
		}, err
	}
	return map[string]string{
		"message": "Product telah terhapus",
	}, nil
}

func UpdateProduct(db *gorm.DB, product models.Product, idProduct int) (*models.Product, error) {
	var tempProduct models.Product
	if err := db.Where("id_outlet = ? AND id = ?", product.IdOutlet, idProduct).First(&tempProduct).Error; err != nil {
		return nil, err
	}
	tempProduct.HargaProduct = product.HargaProduct
	tempProduct.Stock = product.Stock
	tempProduct.NamaProduct = product.NamaProduct
	db.Save(&tempProduct)
	return &tempProduct, nil
}

func UpdatePhotoProduct(db *gorm.DB, filename string, file multipart.File, idUser, idProduct int) (*models.Product, error) {
	var product models.Product
	if err := db.Where("id_user = ? AND id = ?", idUser, idProduct).Find(&product).Error; err != nil {
		return nil, err
	}
	out, err := os.Create("./tmp/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	product.Foto = out.Name()
	db.Save(&product)
	return &product, nil
}
