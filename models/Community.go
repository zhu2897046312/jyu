package models

import (
	"github.com/jyu/utils"
	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name        string `json:"name"`
	OwnerId     string `json:"owner_id"`
	Img         string `json:"img"`
	Description string `json:"description"`
}

func (Community) TableName() string {
	return "community"
}

func CreateCommunity(community Community) *gorm.DB {
	return utils.DB_MySQL.Create(&community)
}

func LoadCommunityList(owner_id string) ([]Community, error) {
	data := make([]Community, 10)
	err := utils.DB_MySQL.Where("owner_id = ?", owner_id).Find(&data).Error
	return data, err
}
