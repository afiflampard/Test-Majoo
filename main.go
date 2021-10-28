package main

import (
	"log"
	"majoo/config"
	"majoo/controllers"
	"majoo/migrations"
	"majoo/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {

	db := config.Connect()
	migrations.Migrate(db)
	controllers.InitiateDB(db)
	r := gin.Default()
	routes := routes.NewUserController(controllers.UserControllers(), controllers.ProductControllers(), controllers.OutletControllers(), controllers.TransaksiControllers())
	r.Use(cors.Default())
	routes.Routes(r)
	log.Fatal(r.Run(":8000"))
}
