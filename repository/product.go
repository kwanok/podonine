package repository

import (
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Seats []Seat
}

func (product *Product) GetId() uint {
	return product.ID
}

func (product *Product) GetSeats() []models.Seat {
	seatRepository := SeatRepository{Db: Gorm}
	return seatRepository.GetSeatsByProductId(product.ID)
}

type ProductRepository struct {
	Db *gorm.DB
}

func (repo *ProductRepository) GetProductById(id uint) models.Product {
	var product Product
	product.ID = id
	result := repo.Db.First(&product)
	if result.Error != nil {
		return nil
	}

	return &product
}
