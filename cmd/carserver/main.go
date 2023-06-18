package main

import (
	// "crud-go/controllers"
	handlers "crud-go/pkg/cars/handlers/car"
	handlerUser "crud-go/pkg/cars/handlers/user"
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	//CAR APIS
    r.POST("/cars", middleware.RequireAuth , handlers.New().CarsCreate)
	r.GET("/cars", middleware.RequireAuth,handlers.New().GetCars)
	r.GET("/cars/:id", middleware.RequireAuth,handlers.New().GetCarByID)
	r.PUT("/cars/:id", middleware.RequireAuth, handlers.New().UpdateCarByID)
	r.DELETE("/cars/:id", middleware.RequireAuth, handlers.New().CarDelete)

	//USER APIS
	r.POST("/signup", handlerUser.Signup)
	r.POST("/login", handlerUser.Login)
	r.GET("/validate", middleware.RequireAuth, handlerUser.Validate)

    r.Run() 
}
 