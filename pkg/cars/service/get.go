package services

import (
	"github.com/gin-gonic/gin"
)

func (s *carService) GetCars(c *gin.Context) {
	cars, err := s.store.GetCars()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get cars"})
		return
	}

	//RESPOND WITH CARS
	c.JSON(200, gin.H{"cars": cars})
}