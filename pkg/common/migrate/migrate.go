package main

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main () {
	initializers.DB.AutoMigrate(&models.Car{})
	initializers.DB.AutoMigrate(&models.User{})
}