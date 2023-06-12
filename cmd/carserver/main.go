package main

import (
	// "crud-go/controllers"
	handlers "crud-go/pkg/cars/handlers/car"
	handlerUser "crud-go/pkg/cars/handlers/user"
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/middleware"
	services "crud-go/pkg/cars/service"
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	//CAR APIS
	carStore := store.GetStore()
	carService := services.NewCreateService(carStore)
	carCreateHandler := handlers.NewCreateHandler(carService)
    r.POST("/cars", middleware.RequireAuth ,carCreateHandler.CarsCreate)
	r.GET("/cars", middleware.RequireAuth,handlers.GetCars)
	r.GET("/cars/:id", middleware.RequireAuth,handlers.GetCarByID)
	r.PUT("/cars/:id", middleware.RequireAuth, handlers.UpdateCarByID)
	r.DELETE("/cars/:id", middleware.RequireAuth, handlers.CarDelete)

	//USER APIS
	r.POST("/signup", handlerUser.Signup)
	r.POST("/login", handlerUser.Login)
	r.GET("/validate", middleware.RequireAuth, handlerUser.Validate)

    r.Run() 
}
 