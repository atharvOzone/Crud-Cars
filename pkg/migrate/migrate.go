package main

import (
	"crud-go/pkg/initializers"
	"crud-go/pkg/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main () {
	initializers.DB.AutoMigrate(&models.Car{})
	initializers.DB.AutoMigrate(&models.User{})
}