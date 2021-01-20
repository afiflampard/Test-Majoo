package main

import (
	"helloworld/config"
	"helloworld/controller"
	"helloworld/migrations"
	"helloworld/routes"
	"log"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {

	var userJwt controller.JWTServices = controller.JWTAuthService()
	var userController controller.UserController = controller.LoginHandler(userJwt)
	var router routes.RouterStruct = routes.NewUserController(userController)

	db := config.Connect()
	migrations.Migrations(db)
	controller.InitiateDB(db)
	r := gin.Default()
	r.Use(cors.Default())
	router.Routes(r)
	log.Fatal(r.Run(":4747"))
}
