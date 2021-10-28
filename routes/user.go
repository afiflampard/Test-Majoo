package routes

import (
	"majoo/controllers"
	"majoo/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Routes(router *gin.Engine)
}

type RouterController struct {
	userService         controllers.UserController
	productService      controllers.ProductController
	outletController    controllers.OutletController
	transaksiController controllers.TransaksiController
}

func NewUserController(userService controllers.UserController, productService controllers.ProductController, outletService controllers.OutletController, transaksiService controllers.TransaksiController) RouterController {
	return RouterController{userService, productService, outletService, transaksiService}
}
func (c *RouterController) Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/login", c.userService.Login)
		v1.POST("/signup", c.userService.SignUp)

		v1.Use(middleware.Authorization())
		v1.GET("/user/:id", c.userService.FindById)
		v1.GET("/user", c.userService.FindAll)
		v1.DELETE("/user/:id", c.userService.Delete)
		v1.PUT("/user/:id", c.userService.Update)
	}

}
