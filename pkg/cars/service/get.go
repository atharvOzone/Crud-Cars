package services

import (
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

type GetService struct {
	store store.Store
}

func NewGetService(store store.Store) *GetService {
	return &GetService{
        store: store,
    }
}

func (s *GetService) GetCars(c *gin.Context) {
	cars, err := s.store.GetCars()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get cars"})
		return
	}

	//RESPOND WITH CARS
	c.JSON(200, gin.H{"cars": cars})
}