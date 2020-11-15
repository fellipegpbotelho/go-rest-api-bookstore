package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api-bookstore/controllers"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
