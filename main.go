package main

import (
	"rest-api-bookstore/models"
	"rest-api-bookstore/routes"
)

func main() {
	models.ConnectDatabase()
	router := routes.CreateRouter()
	err := router.Run()
	if err != nil {
		panic(err.Error())
	}
}