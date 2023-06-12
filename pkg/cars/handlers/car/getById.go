package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"github.com/gin-gonic/gin"
)

func GetCarByID(c *gin.Context) {
	//GET ID FROM URL
	id := c.Param("id")

	//GET CAR BY ID
	var car models.Car
	initializers.DB.First(&car, id)

	//RESPONSD WITH THE DATA
	c.JSON(200, gin.H{
        "car": car,
    })
}
