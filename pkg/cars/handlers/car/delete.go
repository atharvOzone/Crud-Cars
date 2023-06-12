package handlers

import (
	services "crud-go/pkg/cars/service"
	"crud-go/pkg/cars/store"

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

	deleteService := services.NewDeleteService(store.GetStore())
	//Delete car
	err := deleteService.CarDelete(c, id)
	if err!= nil {
        c.JSON(400, gin.H{"error": "Failed to delete the car"})
        return
    }

	//Respond
	c.Status(200)
}