package services

import (
	"crud-go/pkg/cars/models"
	"crud-go/pkg/cars/store"
)

type CarService interface {
	CarsCreate(car models.Car) error
	GetCars(page, pageSize int) (error, []models.Car)
	GetCarByID(id string) (error, models.Car)
	UpdateCarByID(id string, updatedCar models.Car) error
	CarDelete(id string) error
}

type carService struct{
	store store.Store
}

func New() CarService {
	return &carService{
		store: store.GetStore(),
	}
}