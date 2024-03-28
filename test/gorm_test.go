package test

import (
	//"fmt"
	"testing"

	"github.com/jyu/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T){
	dsn := "root:123456@tcp(127.0.0.1:3306)/jyu?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Community{})
}
