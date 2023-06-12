package handlers

import (
	services "crud-go/pkg/cars/service"
	"crud-go/pkg/cars/store"

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
	getService := services.NewGetService(store.GetStore())

    //RESPONSD WITH THE DATA
	getService.GetCars(c)
}
