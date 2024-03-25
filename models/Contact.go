package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerId     uint
	TaraId      uint
	Type        int
	Description string
}

func (table *Contact) TableNanme() string {
	return "contact"
}
