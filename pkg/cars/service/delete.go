package services

import (
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

type DeleteService struct {
	store store.Store
}

func NewDeleteService(store store.Store) *DeleteService {
	return &DeleteService{
		store: store,
	}
}

func (s *DeleteService) CarDelete(c *gin.Context, id string) error {
	//GET ID FROM THE URL
	//DELETE CAR
	err := s.store.CarDelete(id)
	if err!= nil {
        c.JSON(500, gin.H{"error": "Failed to delete car"})
        return err
    }

	//RESPOND WITH SUCCESS
	return nil
}