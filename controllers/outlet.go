package controllers

import "github.com/gin-gonic/gin"

type OutletController interface {
	CreateOutlet(c *gin.Context)
}
