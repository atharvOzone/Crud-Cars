package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"github.com/gin-gonic/gin"
)

// CreateItem GetCars
// @Summary      Get all the cars
// @Description  Get all the cars
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        car body models.Car true "Car"
// @Router       /cars [get]
func GetCars(c *gin.Context) {
	//GET CARS
	var cars []models.Car
	initializers.DB.Find(&cars)

    //RESPONSD WITH THE DATA
	c.JSON(200, gin.H{
		"cars": cars,
	})
}
