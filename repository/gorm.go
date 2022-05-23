package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("sqlite/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm")
	}

	Gorm = db
}
