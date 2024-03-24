package test

import (
	//"fmt"
	"testing"

	"github.com/jyu/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

func TestGorm(t *testing.T){
	dsn := "root:123456@tcp(127.0.0.1:3306)/jyu"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Account = "221110137"

	db.Create(user)

	fmt.Println(db.First(user), 1)

	db.Model(user).Update("password","1234")
}
