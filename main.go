package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-bookstore/controllers"
	"rest-api-bookstore/models"
)

func main() {
	routes := gin.Default()
	models.ConnectDatabase()

	routes.GET("/books", controllers.FindBooks)
	routes.POST("/books", controllers.CreateBook)
	routes.GET("/books/:id", controllers.FindBook)

	err := routes.Run()
	if err != nil {
		panic(err.Error())
	}
}