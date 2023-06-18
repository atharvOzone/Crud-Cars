package services

import (
	"crud-go/pkg/cars/models"
)
func (s *carService) CarsCreate(car models.Car) error{


	//CREATE CAR
	err := s.store.CarsCreate(car)
	if err!= nil {
		return err
    }
	return nil 
}