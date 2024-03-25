package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string
	OwnerId     uint
	Icon        string
	Type        int
	Description string
}

func (table *Group) TableNanme() string {
	return "group"
}
