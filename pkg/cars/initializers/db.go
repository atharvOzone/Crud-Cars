package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ReturnDB() (*gorm.DB, error) {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}