package controllers

import "github.com/gin-gonic/gin"

type TransaksiController interface {
	OrderJual(c *gin.Context)
	OrderBeli(c *gin.Context)
	HistoryPenjualan(c *gin.Context)
	HistoryPembelian(c *gin.Context)
}
