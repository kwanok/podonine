package models

type Product interface {
	GetId() uint
	GetSeats() []Seat
}

type ProductRepository interface {
	GetProductById(id uint) Product
}
