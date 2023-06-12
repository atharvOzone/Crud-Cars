package main

import (
	// "crud-go/controllers"
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/middleware"
	"crud-go/pkg/cars/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	//CAR APIS
    r.POST("/cars", middleware.RequireAuth ,controllers.CarsCreate)
	r.GET("/cars", middleware.RequireAuth,controllers.GetCars)
	r.GET("/cars/:id", middleware.RequireAuth,controllers.GetCarByID)
	r.PUT("/cars/:id", middleware.RequireAuth, controllers.UpdateCarByID)
	r.DELETE("/cars/:id", middleware.RequireAuth, controllers.CarDelete)

	//USER APIS
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

    r.Run() 
}
 