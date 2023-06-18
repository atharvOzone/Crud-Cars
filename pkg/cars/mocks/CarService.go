// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	models "crud-go/pkg/cars/models"
)

// CarService is an autogenerated mock type for the CarService type
type CarService struct {
	mock.Mock
}

// CarDelete provides a mock function with given fields: id
func (_m *CarService) CarDelete(id string) {
	_m.Called(id)
}

// CarsCreate provides a mock function with given fields: c, car
func (_m *CarService) CarsCreate(c *gin.Context, car *models.Car) {
	_m.Called(c, car)
}

// GetCarByID provides a mock function with given fields: c
func (_m *CarService) GetCarByID(c *gin.Context) {
	_m.Called(c)
}

// GetCars provides a mock function with given fields: c
func (_m *CarService) GetCars(c *gin.Context) {
	_m.Called(c)
}

// UpdateCarByID provides a mock function with given fields: c
func (_m *CarService) UpdateCarByID(c *gin.Context) {
	_m.Called(c)
}

// NewCarService creates a new instance of CarService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCarService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CarService {
	mock := &CarService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}