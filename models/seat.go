package models

type Seat interface {
}

type SeatRepository interface {
	GetSeatsByProductId(productId uint) []Seat
}
