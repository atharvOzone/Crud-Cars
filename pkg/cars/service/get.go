package services

import (
	"crud-go/pkg/cars/models"
)

func (s *carService) GetCars(page, pageSize int) (error, []models.Car) {
	// cars, err := s.store.GetCars()
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to get cars"})
	// 	return
	// }

	// //RESPOND WITH CARS
	// c.JSON(200, gin.H{"cars": cars})
	
	offset := (page	- 1) * pageSize
	err, car := s.store.GetCars(page, pageSize, offset)
	if err != nil {
		return err, nil
	}
	return nil, car
}