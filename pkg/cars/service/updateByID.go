package services

import (
	"crud-go/pkg/cars/models"
	"crud-go/pkg/cars/store"
)

func (s *carService) UpdateCarByID(id string, updatedCar models.Car) error {
	err, car := s.store.GetCarByID(id)
	if err!=nil {
		return err
	}
	err = store.GetStore().UpdateCarByID(car, updatedCar)
	if err != nil {
		return err
	}
	return nil
}