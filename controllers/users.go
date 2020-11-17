package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-api-bookstore/models"
	"time"
)

type RegisterUserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Age uint8 `json:"age" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserOutput struct {
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age uint8 `json:"age"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Register a new user
func RegisterUser(context *gin.Context) {
	var input RegisterUserInput

	inputError := context.ShouldBindJSON(&input)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": inputError.Error()})
		return
	}

	var existingUser models.User
	models.DB.Where("email = ?", input.Email).First(&existingUser)

	if existingUser.ID != 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "E-mail already exists."})
		return
	}

	hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if hashError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": hashError.Error()})
		return
	}

	userToCreate := models.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Age: input.Age,
		Email: input.Email,
		Password: string(hashedPassword),
	}
	queryUserCreationError := models.DB.Create(&userToCreate).Error
	if queryUserCreationError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": queryUserCreationError.Error()})
		return
	}

	output := RegisterUserOutput{
		ID: userToCreate.ID,
		FirstName: userToCreate.FirstName,
		LastName: userToCreate.LastName,
		Age: userToCreate.Age,
		Email: userToCreate.Email,
		CreatedAt: userToCreate.CreatedAt,
		UpdatedAt: userToCreate.UpdatedAt,
	}
	context.JSON(http.StatusOK, gin.H{"data": output})
}
