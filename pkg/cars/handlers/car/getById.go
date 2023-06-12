package handlers

import (
	services "crud-go/pkg/cars/service"
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

// CreateItem GetCars
// @Summary      Get car by id
// @Description  Get car by id
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        car body models.Car true "Car"
// @Router       /cars/:id [get]
func GetCarByID(c *gin.Context) {
	//GET ID FROM URL
	getService := services.NewGetServiceID(store.GetStore())

	//RESPONSD WITH THE DATA
	getService.GetCarByID(c)
}
