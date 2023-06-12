package handlers

import (
	services "crud-go/pkg/cars/service"
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

// CreateItem UpdateCarByID
// @Summary      updates the car by given id
// @Description  updates the car by given id
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        car body models.Car true "Car"
// @Router       /cars/:id [put]
func UpdateCarByID(c *gin.Context) {
	//GET ID FROM URL
	updateService := services.NewUpdateService(store.GetStore())
	updateService.UpdateCarByID(c)
	//RESPOND WITH THE DATA

}