package controllers

import (
	"fmt"
	"majoo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Outlet struct{}

func OutletControllers() OutletController {
	return Outlet{}
}

func (ctx Outlet) CreateOutlet(c *gin.Context) {
	idUser := int(GetID(c))
	fmt.Println("ID user", idUser)
	resp, err := service.CreateOutlet(GetDB(), idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request tidak sesuai")
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
