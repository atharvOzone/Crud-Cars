package store

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"gorm.io/gorm"
)

type Store interface {
	CarsCreate(car models.Car) error
	CarDelete(id string) error
	GetCars(page, pageSize, offset int) (error, []models.Car)
	GetCarByID(id string) (error, models.Car)
	UpdateCarByID(initialCar, updatedCar models.Car) error
}

type DBCarStore struct {
	DB *gorm.DB
}


var carStore Store


func NewStore() Store {
	db, err := initializers.ReturnDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	return &DBCarStore{
		DB: db,
	}
}

func GetStore() Store {
	if carStore == nil {
        carStore = NewStore()
    }
    return carStore
}


func (s *DBCarStore) CarsCreate(car models.Car) error {
	return s.DB.Create(car).Error
}

func (s *DBCarStore) CarDelete(id string) error {
	return s.DB.Delete(&models.Car{}, id).Error
}

func (s *DBCarStore) GetCars(page, pageSize, offset int) (error, []models.Car) {
	var cars []models.Car
	return s.DB.Offset(offset).Limit(pageSize).Find(&cars).Error, cars
}

func (s *DBCarStore) GetCarByID(id string) (error, models.Car) {
	var car models.Car
	return s.DB.First(&car, id).Error, car
}

func (s *DBCarStore) UpdateCarByID(initialCar, updatedCar models.Car) error {
	return s.DB.Model(&initialCar).Where("car_id = ?", initialCar.CAR_ID).Updates(updatedCar).Error
}