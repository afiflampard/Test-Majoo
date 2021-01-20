package routes

import (
	"helloworld/controller"
	"helloworld/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Routes(router *gin.Engine)
}

type RouterStruct struct {
	userService controller.UserController
}

func NewUserController(userController controller.UserController) RouterStruct {
	return RouterStruct{userController}
}

func (c *RouterStruct) Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/login", c.userService.Login)
		v1.POST("/signup", c.userService.SignUp)

		v1.Use(middleware.Authorization())
		v1.GET("/user/:id", c.userService.FindById)
		v1.GET("/user", c.userService.FindAll)
		v1.PUT("/user/:id", c.userService.Update)
		v1.DELETE("/user/:id", c.userService.Delete)
		v1.PUT("/user/:id/photo", c.userService.UpdatePhoto)
	}

}
