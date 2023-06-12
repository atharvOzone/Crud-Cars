package handlers

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"github.com/gin-gonic/gin"
)
// CreateItem UpdateCarByID
// @Summary      updates the car by given id
// @Description  updates the car by given id
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        car body models.Car true "Car"
// @Router       /cars/:id [put]
func UpdateCarByID(c *gin.Context) {
	//GET ID FROM URL
	id := c.Param("id")

	//GET DATA FROM REQUEST
	var carContent struct {
		ID uint
        COMPANY_NAME string
        CAR_NAME string
        YEAR int64
        PRICE float64
	}

	c.Bind(&carContent)

	//FIND THE UPDATING POST
	var car models.Car
	initializers.DB.First(&car, id)

	//UPDATE DATA
	initializers.DB.Model(&car).Updates(models.Car{
		ID: carContent.ID,
		COMPANY_NAME: carContent.COMPANY_NAME,
		CAR_NAME: carContent.CAR_NAME,
        YEAR: carContent.YEAR,
        PRICE: carContent.PRICE,
	})

	//RESPOND WITH THE DATA
	c.JSON(200, gin.H{
		"car": car,
	})

}