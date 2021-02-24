package controller

import (
	"fmt"
	"helloworld/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ActivityImpl struct {
	jwtServices JWTServices
}

type RequestPinjam struct {
	JudulBuku      string    `json:"judul_buku"`
	TanggalKembali time.Time `json:"tanggal_kembali"`
}

type SuccessPinjam struct {
	Kode    uint   `json:"status"`
	Message string `json:"message"`
}

func NewActivityController(jwtservices JWTServices) ActivityController {
	return &ActivityImpl{jwtservices}
}

func (controller *ActivityImpl) PinjamBuku(c *gin.Context) {
	idMember := c.Param("id")
	idPetugas := c.Query("idPetugas")
	// fmt.Println(idMember)
	// fmt.Println(idPetugas)
	var buku models.Buku
	var req RequestPinjam
	var petugas models.User
	var member models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Harus JSON ya")
	}
	fmt.Println(req)
	if err := GetDB().Where("judul_buku = ?", req.JudulBuku).First(&buku).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	}
	if err := GetDB().Model(&petugas).Preload("Role").Find(&petugas, idPetugas).Error; err != nil {
		Error(c, 404, "Petugas Not Found")
	}

	if err := GetDB().Model(&member).Preload("Role").Find(&member, idMember).Error; err != nil {
		Error(c, 404, "Member Not Found")
	}
	fmt.Println(petugas)
	fmt.Println(member)

	if strings.ToLower(petugas.Role.Role) == "petugas" {
		pinjam := models.Order{
			TanggalPeminjaman: time.Now(),
			TanggalKembali:    req.TanggalKembali,
			IDPetugas:         petugas.ID,
			IDUser:            member.ID,
			NoState:           1,
		}
		err := GetDB().Debug().Create(&pinjam).Error
		if err != nil {
			c.JSON(401, &ErrorResponse{
				Error: err,
			})
		} else {
			orderDetail := models.OrderDetail{
				IDOrder: pinjam.ID,
				IDBuku:  buku.ID,
			}
			err := GetDB().Debug().Create(&orderDetail).Error
			if err != nil {
				c.JSON(401, &ErrorResponse{
					Error: err,
				})
			}
			buku.Stok = buku.Stok - 1
			GetDB().Save(&buku)
			history := models.History{
				IDBuku:  buku.ID,
				IDOrder: pinjam.ID,
				NoState: pinjam.NoState,
			}
			err = GetDB().Debug().Create(&history).Error
			if err != nil {
				c.JSON(401, &ErrorResponse{
					Error: err,
				})
			} else {
				c.JSON(200, &SuccessPinjam{
					Kode:    200,
					Message: "Buku Sudah Dipinjam",
				})
			}
		}
	}

}

func (controller *ActivityImpl) KembaliBuku(c *gin.Context) {
	idMember := c.Param("id")
	idPetugas := c.Query("idPetugas")
	var buku models.Buku
	var req RequestPinjam
	var petugas models.User
	var member models.User

	var orderDetail models.OrderDetail
	var order models.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Harus JSON ya")
	}

	if err := GetDB().Where("judul_buku = ?", req.JudulBuku).First(&buku).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	}
	if err := GetDB().Model(&petugas).Preload("Role").Find(&petugas, idPetugas).Error; err != nil {
		Error(c, 404, "Petugas Not Found")
	}

	if err := GetDB().Model(&member).Preload("Role").Find(&member, idMember).Error; err != nil {
		Error(c, 404, "Member Not Found")
	}
	if err := GetDB().Model(&orderDetail).Where("buku_id = ?", buku.ID).Preload("Order").Find(&orderDetail).Error; err != nil {
		Error(c, 404, "Buku Not Found")
	}
	if err := GetDB().Where("id = ?", orderDetail.IDOrder).First(&order).Error; err != nil {
		Error(c, 404, "Order Not Found")
	}
	if strings.ToLower(petugas.Role.Role) == "petugas" {
		orderDetail.Order.NoState = 2
		order.NoState = 2
		buku.Stok = buku.Stok + 1
		GetDB().Save(&order)
		GetDB().Save(&buku)
		history := models.History{
			IDBuku:  buku.ID,
			IDOrder: orderDetail.Order.ID,
			NoState: orderDetail.Order.NoState,
		}
		err := GetDB().Debug().Create(&history).Error
		if err != nil {
			c.JSON(401, &ErrorResponse{
				Error: err,
			})
		}
		c.JSON(200, "Data Telah Terupdate")

	}
}
