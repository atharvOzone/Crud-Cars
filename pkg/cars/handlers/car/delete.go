package handlers

import (
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
func (h *Handler) CarDelete(c *gin.Context) {
	//GET THE ID FROM URL
	id := c.Param("id")

	// deleteService := services.NewCarService(store.GetStore())
	//Delete car
	err := h.Service.CarDelete(id)
	if err!= nil {
        c.JSON(400, gin.H{"error": "Failed to delete the car"})
        return
    }

	//Respond
	c.JSON(200, gin.H{
		"status": "car deleted successfully",
	})
}