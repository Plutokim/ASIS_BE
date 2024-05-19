package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dbstring string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbstring), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
