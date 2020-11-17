package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api-bookstore/controllers"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	users.POST("/register", controllers.RegisterUser)

	books := router.Group("/books")
	books.GET("/", controllers.FindBooks)
	books.POST("/", controllers.CreateBook)
	books.GET("/:id", controllers.FindBook)
	books.PATCH("/:id", controllers.UpdateBook)
	books.DELETE("/:id", controllers.DeleteBook)

	return router
}
