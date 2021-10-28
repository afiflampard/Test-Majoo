package controllers

import "github.com/gin-gonic/gin"

type ProductController interface {
	CreateProduct(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UpdatePhoto(c *gin.Context)
}
