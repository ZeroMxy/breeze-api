package main

import (
	"breeze-api/config"
	"breeze-api/database"
	"breeze-api/internal/router"
	"breeze-api/pkg/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {

	// 实例化
	app := fiber.New()

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := config.Mysql.Username + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Host + ":" + strconv.Itoa(config.Mysql.Port) + ")/" + config.Mysql.Database + "?charset=utf8mb4&parseTime=True"

	// 初始化数据访问层
	db.Init(&db.DBConfig{
		GormConfig: &db.GormConfig{
			Dialector: mysql.Open(dsn),
			Opts: &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			},
		},
	})

	// 根目录
	app.Static("/", "./web")

	// 初始化数据库
	database.Init()

	// 注册路由
	router.AdminRouter(app)

	// 恢复
	app.Use(recover.New())

	app.Listen(config.App.Host + ":" + strconv.Itoa(config.App.Port))
}
