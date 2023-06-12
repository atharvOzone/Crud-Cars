package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID uint
	COMPANY_NAME string
	CAR_NAME string
	YEAR int64
	PRICE float64
}