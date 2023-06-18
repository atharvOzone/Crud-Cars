package handlers

import (
	"crud-go/pkg/cars/models"
	"net/http"

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
func (h *Handler) UpdateCarByID(c *gin.Context) {
	//GET ID FROM URL
	// updateService := services.NewCarService(store.GetStore())
	// updateService.UpdateCarByID(c)
	// //RESPOND WITH THE DATA

	id := c.Param("car_id")
	req := models.CreateCarRequest{}
	err := c.Bind(&req)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}
	car := req.ToCar()

	err = h.Service.UpdateCarByID(id, car)

	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated Successfully",
	})
	

}