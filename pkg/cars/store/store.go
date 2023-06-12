package store

import (
	"crud-go/pkg/cars/initializers"
	"crud-go/pkg/cars/models"

	"gorm.io/gorm"
)

type Store interface {
	CarsCreate(car *models.Car) error
	CarDelete(id string) error
	GetCars() ([]models.Car, error)
	GetCarByID(id string) (*models.Car, error)
	UpdateCarByID(car *models.Car) error
}

type DBCarStore struct {
	DB *gorm.DB
}

func NewStore() Store {
	db, err := initializers.ReturnDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	return &DBCarStore{
		DB: db,
	}
}

var carStore Store

func GetStore() Store {
	if carStore == nil {
        carStore = NewStore()
    }
    return carStore
}


func (s *DBCarStore) CarsCreate(car *models.Car) error {
	result := s.DB.Create(car)
	if result.Error!= nil {
        return result.Error
    }
	return nil
}

func (s *DBCarStore) CarDelete(id string) error {
	result := s.DB.Delete(&models.Car{},id)
	if result.Error!= nil {
        return result.Error
    }
	return nil
}

func (s *DBCarStore) GetCars() ([]models.Car, error) {
	var cars []models.Car
    result := s.DB.Find(&cars)
    if result.Error!= nil {
        return nil, result.Error
    }
    return cars, nil
}

func (s *DBCarStore) GetCarByID(id string) (*models.Car, error) {
	var car models.Car
	result := s.DB.First(&car, id)
	if result.Error!= nil {
        return nil, result.Error
    }
	return &car, nil
}

func (s *DBCarStore) UpdateCarByID(car *models.Car) error {
	result := s.DB.Model(&car).Updates(car)
	if result.Error!= nil {
        return result.Error
    }
	return nil
}