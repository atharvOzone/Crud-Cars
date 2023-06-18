package handlers

import (
	"net/http"
	"strconv"

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
func (h *Handler) GetCars(c *gin.Context) {
	pagestr := c.Query("page")
	page, err := strconv.Atoi(pagestr)
	if err!=nil{
		page = 1
	}
	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err!=nil{
		pageSize = 10
	}

	err, car := h.Service.GetCars(page, pageSize)
	if err!=nil {
		c.JSON(500, gin.H{"error": "Failed to load required data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Car": car,
	})
}
