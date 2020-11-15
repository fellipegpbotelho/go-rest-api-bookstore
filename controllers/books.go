package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-bookstore/models"
)

type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// Get all books
func FindBooks(context *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	context.JSON(http.StatusOK, gin.H{"data": books})
}

// Create new book
func CreateBook(context *gin.Context) {
	var input CreateBookInput
	inputErr := context.ShouldBindJSON(&input)

	if inputErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": inputErr.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	context.JSON(http.StatusOK, gin.H{"data": book})
}