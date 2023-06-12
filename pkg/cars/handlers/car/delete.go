package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"github.com/gin-gonic/gin"
)

// CreateItem CarDelete
// @Summary      Delete a Car
// @Description  Delete a car by Id
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        car body models.Car true "Car"
// @Router       /cars/:id [delete]
func CarDelete(c *gin.Context) {
	//GET THE ID FROM URL
	id := c.Param("id")

	//Delete car
	initializers.DB.Delete(&models.Car{}, id)

	//Respond
	c.Status(200)
}