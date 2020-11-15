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

type UpdateBookInput struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

// Get all books
func FindBooks(context *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)
	context.JSON(http.StatusOK, gin.H{"data": books})
}

// Find a book
func FindBook(context *gin.Context) {
	var book models.Book

	queryErr := models.DB.Where("id = ?", context.Param("id")).First(&book).Error
	if queryErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Book not found."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": book})
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

// Update a book
func UpdateBook(context *gin.Context) {
	var book models.Book

	queryErr := models.DB.Where("id = ?", context.Param("id")).First(&book).Error
	if queryErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Book not found."})
		return
	}

	var input UpdateBookInput

	inputErr := context.ShouldBindJSON(&input)
	if inputErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": inputErr.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete a book
func DeleteBook(context *gin.Context) {
	var book models.Book

	queryErr := models.DB.Where("id = ?", context.Param("id")).First(&book).Error
	if queryErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Book not found."})
		return
	}

	models.DB.Delete(&book)

	context.JSON(http.StatusOK, gin.H{"data": true})
}