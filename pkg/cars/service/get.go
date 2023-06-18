package services

import (
	"crud-go/pkg/cars/models"
)

func (s *carService) GetCars(page, pageSize int) (error, []models.Car) {
	offset := (page	- 1) * pageSize
	err, car := s.store.GetCars(page, pageSize, offset)
	if err != nil {
		return err, nil
	}
	return nil, car
}