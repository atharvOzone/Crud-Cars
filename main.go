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
    r.POST("/cars", controllers.CarsCreate)
    r.Run() 
}
 