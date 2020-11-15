package main

import (
	"rest-api-bookstore/models"
	"rest-api-bookstore/routes"
)

func main() {
	models.ConnectDatabase()
	err := routes.CreateRouter().Run()
	if err != nil {
		panic(err.Error())
	}
}