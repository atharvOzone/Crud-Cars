package main

import (
	"crud-go/controllers"
	"crud-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	//CAR APIS
    r.POST("/cars", controllers.CarsCreate)
	r.GET("/cars", controllers.GetCars)
	r.GET("/cars/:id", controllers.GetCarByID)
	r.PUT("/cars/:id", controllers.UpdateCarByID)
	r.DELETE("/cars/:id", controllers.CarDelete)

	//USER APIS
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

    r.Run() 
}
 