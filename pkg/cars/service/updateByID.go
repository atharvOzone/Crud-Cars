package services

import (
	"crud-go/pkg/cars/models"
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

type UpdateService struct {
	store store.Store
}

func NewUpdateService(store store.Store) *UpdateService {
	return &UpdateService{
        store: store,
    }
}

func (s *UpdateService) UpdateCarByID(c *gin.Context) {
	//GET ID FROM THE URL
	id := c.Param("id")

	//GET DATA FROM REQUEST BODY
	var carContent models.Car
	if err := c.BindJSON(&carContent); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	//GET THE UPDATED CAR FROM STORE
	car, err := s.store.GetCarByID(id)
	if err!= nil {
        c.JSON(500, gin.H{"error": "Faild to get the updated car"})
        return
    }

	//UPDATED DATA
	car.ID = carContent.ID
	car.COMPANY_NAME = carContent.COMPANY_NAME
	car.CAR_NAME = carContent.CAR_NAME
	car.YEAR = carContent.YEAR
	car.PRICE = carContent.PRICE

	//SAVE THE UPDATED CAR IN STORE
	err = s.store.UpdateCarByID(car)
    if err!= nil {
        c.JSON(500, gin.H{"error": "Faild to update the car"})
        return
    }

    c.JSON(200, gin.H{"car": car})
}