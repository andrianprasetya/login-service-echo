package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"login-service/entity"
	"os"
)

var dataBase *gorm.DB

func New() *gorm.DB {
	connection :=
			"host=" + os.Getenv("DB_HOST") +
			" port=" + os.Getenv("DB_PORT") +
			" user=" + os.Getenv("DB_USERNAME") +
			" dbname=" + os.Getenv("DB_NAME") +
			" password=" + os.Getenv("DB_PASSWORD") +
			" sslmode=" + os.Getenv("DB_SSL") +
			" TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		fmt.Println("Error DB: ", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(3)
	dataBase = db
	return db
}

// GetLinkDb :
func GetLinkDb() *gorm.DB {
	return dataBase
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		entity.User{},
	)
}
