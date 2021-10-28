package controllers

import (
	"majoo/entities"
	"majoo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaksi struct{}

func TransaksiControllers() TransaksiController {
	return Transaksi{}
}

func (ctx Transaksi) OrderJual(c *gin.Context) {
	idUser := int(GetID(c))
	request := entities.OrderRequest{}
	if err := c.ShouldBindJSON(request).Error; err != nil {
		c.JSON(http.StatusBadRequest, "Request not valid")
		return
	}
	request.IDPembeli = uint(idUser)
	resp, err := service.TransaksiJual(GetDB(), request)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Invalid transaksi")
		return
	}
	c.JSON(http.StatusAccepted, resp)

}

func (ctx Transaksi) HistoryPenjualan(c *gin.Context) {
	idOutlet, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}
	resp, err := service.HistoryPenjualan(GetDB(), idOutlet)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Inventory Kosong")
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx Transaksi) OrderBeli(c *gin.Context) {
	idUser := int(GetID(c))
	request := entities.OrderRequest{}
	if err := c.ShouldBindJSON(request).Error; err != nil {
		c.JSON(http.StatusBadRequest, "Request not valid")
		return
	}
	request.IDPembeli = uint(idUser)
	resp, err := service.TransaksiBeli(GetDB(), request)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Invalid transaksi")
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
func (ctx Transaksi) HistoryPembelian(c *gin.Context) {
	idOutlet, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}
	resp, err := service.HistoryPenjualan(GetDB(), idOutlet)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Inventory Kosong")
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
