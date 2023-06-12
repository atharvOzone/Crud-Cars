package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"github.com/gin-gonic/gin"
)

func CarDelete(c *gin.Context) {
	//GET THE ID FROM URL
	id := c.Param("id")

	//Delete car
	initializers.DB.Delete(&models.Car{}, id)

	//Respond
	c.Status(200)
}