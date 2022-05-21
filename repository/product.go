package repository

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Seats []Seat
}
