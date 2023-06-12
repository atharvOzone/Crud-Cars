package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	//GET CARS
	var cars []models.Car
	initializers.DB.Find(&cars)

    //RESPONSD WITH THE DATA
	c.JSON(200, gin.H{
		"cars": cars,
	})
}
