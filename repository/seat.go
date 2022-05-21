package repository

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	ProductId uint
}

func (s *Seat) GetId() uint {
	return s.ID
}
