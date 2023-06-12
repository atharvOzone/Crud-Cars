package services

import (
	"crud-go/pkg/cars/models"
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

type CreateService struct {
	store store.Store
}

func NewCreateService(store store.Store) *CreateService {
	return &CreateService{
		store: store,
	}
}

func (s *CreateService) CarsCreate(c *gin.Context, car *models.Car){


	//CREATE CAR
	err := s.store.CarsCreate(car)
	if err!= nil {
        c.JSON(500, gin.H{"error": "Failed to create car"})
        return
    }

	//RESPOND WITH THE CREATED CAR
	c.JSON(200, gin.H{
		"car": car,
	})
}