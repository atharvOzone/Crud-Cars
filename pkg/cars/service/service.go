package services

import (
	"crud-go/pkg/cars/models"
	"crud-go/pkg/cars/store"

	"github.com/gin-gonic/gin"
)

type CarService interface {
	CarsCreate(c *gin.Context, car *models.Car)
	GetCars(c *gin.Context)
	GetCarByID(c *gin.Context)
	UpdateCarByID(c *gin.Context)
	CarDelete(id string)
}

type carService struct{
	store store.Store
}

func NewCarService(store store.Store) *carService {
	return &carService{
		store: store,
	}
}