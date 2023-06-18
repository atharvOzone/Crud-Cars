package services

import "crud-go/pkg/cars/models"

func (s *carService) GetCarByID(id string) (error, models.Car){
	err, car := s.store.GetCarByID(id)
	if err!=nil{
		return err, models.Car{}
	}
	return nil, car
}