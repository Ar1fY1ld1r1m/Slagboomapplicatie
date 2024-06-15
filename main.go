package main

import (
	"log"
	"slagboomapplicatie_code/api"
	"slagboomapplicatie_code/database"
	"slagboomapplicatie_code/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Slagboomapplicatie API
// @version 1
// @description API voor de slagboomapplicatie

// @host localhost:8080
// @BasePath /
// @schemes http
// @produces json
// @consumes json

func main() {

	err := database.Connect()
	if err != nil {
		log.Fatalf("database is ontoegankelijk: %v", err)
	}

	r := gin.Default()
	docs.SwaggerInfo.Title = "Slagboomapplicatie API"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/check", func(c *gin.Context) {
		kenteken := c.Query("licenseplate")
		api.CheckKenteken(kenteken, c)
	})

	r.Run()

}
