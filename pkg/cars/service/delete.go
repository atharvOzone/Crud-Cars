package services

import (
	"github.com/gin-gonic/gin"
)

func (s *carService) CarDelete(c *gin.Context, id string) error {
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