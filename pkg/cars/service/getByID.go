package services

import (
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

type GetServiceID struct {
	store store.Store
}

func NewGetServiceID(store store.Store) *GetServiceID {
	return &GetServiceID{
        store: store,
    }
}

func (s *GetServiceID) GetCarByID(c *gin.Context) {
	//GET ID FROM URL
	id := c.Param("id")

	//GET CAR BY ID FROM STORE
	car, err := s.store.GetCarByID(id)
	if err!= nil {
        c.JSON(500, gin.H{"error": "Failed to get car"})
        return
    }

	//REPOND WITH CAR
	c.JSON(200, gin.H{"car": car})
}