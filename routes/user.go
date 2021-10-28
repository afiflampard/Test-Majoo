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
	v2 := router.Group("/v1")
	{
		v3 := v2.Group("/outlet")
		{
			v3.Use(middleware.Authorization())
			v3.POST("/", c.outletController.CreateOutlet)
		}
	}
	v4 := router.Group("/v1")
	{
		v5 := v4.Group("/product")
		{
			v5.Use(middleware.Authorization())
			v5.POST("/:id", c.productService.CreateProduct)
			v5.GET("/:id", c.productService.FindById)
			v5.GET("/all/:id", c.productService.FindAll)
			v5.PUT("/:id", c.productService.Update)
			v5.DELETE("/:id", c.productService.Delete)
			v5.PUT("/updatephoto/:id", c.productService.UpdatePhoto)
		}
	}
	v6 := router.Group("/v1")
	{
		v7 := v6.Group("/transaksi")
		{
			v7.Use(middleware.Authorization())
			v7.POST("/jual", c.transaksiController.OrderJual)
			v7.GET("/historyjual/:id", c.transaksiController.HistoryPenjualan)
			v7.GET("/historybeli/:id", c.transaksiController.HistoryPembelian)
			v7.POST("/beli", c.transaksiController.OrderBeli)
		}
	}
}
