package database

import (
	config "studioj/boilerplate_go/configs"
	"studioj/boilerplate_go/internal/models"

	"gorm.io/driver/sqlite" // MySQL 대신 SQLite 드라이버 사용
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() {
	DB = ConnectDB(config.DATABASE_URL) // 괄호 제거
}

func ConnectDB(url string) *gorm.DB {
	gorm_config := gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}}
	db, err := gorm.Open(sqlite.Open(url), &gorm_config) // MySQL 대신 SQLite 사용
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func AutoMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Todo{},
	)
}
