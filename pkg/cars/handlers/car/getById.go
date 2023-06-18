package handlers

import (
	"net/http"

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
func (h *Handler) GetCarByID(c *gin.Context) {
	//GET ID FROM URL
	// id := c.Param("id")
	// getService := services.NewCarService(store.GetStore())

	// //RESPONSD WITH THE DATA
	// getService.GetCarByID(c)

	id := c.Param("id")
	err, car := h.Service.GetCarByID(id)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to recover the movie",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Car": car,
	})

}
