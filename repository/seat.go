package repository

import (
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model
	ProductId uint
}

type SeatRepository struct {
	Db *gorm.DB
}

func (s *Seat) GetId() uint {
	return s.ID
}

func (repo *SeatRepository) GetSeatsByProductId(productId uint) []models.Seat {
	var _seats []Seat
	repo.Db.Model(Seat{ProductId: productId}).Find(&_seats)

	var seats []models.Seat

	for _, seat := range _seats {
		seats = append(seats, seat)
	}

	return seats
}
