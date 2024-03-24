package utils

import (
	"fmt"
	"os"
	"time"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig(){
	viper.SetConfigName("app")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil{
		fmt.Println(err)
	}
}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout,"\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/jyu?charset=utf8&parseTime=True&loc=Local"), 
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}