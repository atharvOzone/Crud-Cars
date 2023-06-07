package controllers

import (
	"crud-go/initializers"
	"crud-go/models"

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

func GetCars(c *gin.Context) {
	//GET CARS
	var cars []models.Car
	initializers.DB.Find(&cars)

    //RESPONSD WITH THE DATA
	c.JSON(200, gin.H{
		"cars": cars,
	})
}

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