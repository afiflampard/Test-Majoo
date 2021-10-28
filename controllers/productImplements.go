package controllers

import (
	"fmt"
	"majoo/entities"
	"majoo/models"
	"majoo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct{}

func ProductControllers() ProductController {
	return Product{}
}

func (ctx Product) CreateProduct(c *gin.Context) {

	newProduct := &entities.RequestProduct{}
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		fmt.Println("Errornya ", err)
		c.JSON(http.StatusBadRequest, "Request Tidak Valid")
		return
	}
	idUser := GetID(c)
	id := c.Param("id")
	IdOutlet, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Request Tidak Valid")
	}

	resp, err := service.CreateProduct(GetDB(), *newProduct, IdOutlet, int(idUser))
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"message": "You cannot add Product",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)

}
func (ctx Product) FindById(c *gin.Context) {
	id := c.Param("id")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request tidak valid")
		return
	}
	resp, err := service.GetProductByID(GetDB(), idProduct)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Product Not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx Product) FindAll(c *gin.Context) {
	id := c.Param("id")
	idOutlet, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request Not Valid")
		return
	}
	resp, err := service.GetAllProduct(GetDB(), idOutlet)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Product not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx Product) Update(c *gin.Context) {
	product := models.Product{}
	id := c.Param("id")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request Not Valid")
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, "Request not valid")
		return
	}
	resp, err := service.UpdateProduct(GetDB(), product, idProduct)
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"message": "You cannot update product",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx Product) Delete(c *gin.Context) {
	id := c.Param("id")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request not valid")
		return
	}
	resp, err := service.DeleteProduct(GetDB(), idProduct)
	if err != nil {
		c.JSON(http.StatusBadGateway, resp)
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx Product) UpdatePhoto(c *gin.Context) {
	idUser := int(GetID(c))
	idProduct, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request tidak valid")
		return
	}
	file, header, err := c.Request.FormFile("photo")
	filename := header.Filename
	resp, err := service.UpdatePhotoProduct(GetDB(), filename, file, idUser, idProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request tidak bisa disimpan")
		return
	}
	c.JSON(http.StatusAccepted, resp)

}
