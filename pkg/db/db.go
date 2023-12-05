package db

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BD struct{}

type DBConfig struct {
	GormConfig  *GormConfig
	RedisConfig *RedisConfig
}

type GormConfig struct {
	Dialector gorm.Dialector
	Opts      gorm.Option
}

type RedisConfig struct {
	Host     string
	Port     int
	Database int
	Password string
}

var (
	GormClient  *gorm.DB
	RedisClient *redis.Client
)

// 初始化数据访问层
func Init(config *DBConfig) {

	// 初始化 Gorm
	if config.GormConfig != nil {
		initGorm(config.GormConfig.Dialector, config.GormConfig.Opts)
	}

	// 初始化 Redis
	if config.RedisConfig != nil {
		initRedis(&redis.Options{
			Addr:     config.RedisConfig.Host + ":" + strconv.Itoa(config.RedisConfig.Port),
			Password: config.RedisConfig.Password,
			DB:       config.RedisConfig.Database,
		})
	}

	return
}

// 初始化 Gorm
func initGorm(dialector gorm.Dialector, opts gorm.Option) {

	var err error

	GormClient, err = gorm.Open(dialector, opts)
	if err != nil {
		panic(err)
	}

	sqlDB, err := GormClient.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	return
}

// 初始化 Redis
func initRedis(opts *redis.Options) {

	RedisClient = redis.NewClient(opts)

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return
}
