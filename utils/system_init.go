package utils

import (
	"log"
	"os"
	"time"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB_MySQL *gorm.DB
	DB_Redis *redis.Client
)

func InitConfig(path  string) {
	viper.SetConfigName("app")
	if path == "" {
		path = "config"
	}
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	log.Println(viper.GetString("mysql.dns"))

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	DB_MySQL = db
}

func InitRedis() {
	DB_Redis = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
}

const (
	PublishKey = "websocket"

)

// Publish 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) (error){
	log.Println("Publish: " + msg)
	err := DB_Redis.Publish(ctx, channel, msg).Err()
	return err
}
// Subscribe 订阅消息
func Subscribe(ctx context.Context, channel string) (string, error){
	sub:= DB_Redis.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	log.Println("Subscribe: " , msg.Payload)
	return msg.Payload, err
}


