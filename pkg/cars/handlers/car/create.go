package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"github.com/gin-gonic/gin"
)

func CarsCreate(c *gin.Context) {
	//Get data off requ body
	var carContent struct {
		ID uint
		COMPANY_NAME string
		CAR_NAME string
		YEAR int64
		PRICE float64
	}

	c.Bind(&carContent)

	//Create cars
	car := models.Car{ID: carContent.ID, COMPANY_NAME: carContent.COMPANY_NAME, CAR_NAME: carContent.CAR_NAME, YEAR: carContent.YEAR, PRICE: carContent.PRICE}
	result := initializers.DB.Create(&car)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"car": car,
	})
}