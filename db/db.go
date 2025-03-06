package db

import (
	"bmsp-backend-service/models"
	"bmsp-backend-service/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBPg *gorm.DB

func InitDB() *gorm.DB {
	dsn := "user=bmsp_admin password=GyBUPbRngIY8Jth dbname=bmsp host=172.236.141.236 port=5432 sslmode=disable"
	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get DB object:", err)
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(25)

	sqlDB.SetConnMaxLifetime(60)

	fmt.Println("Successfully connected to the database!")

	DBPg = DB

	return DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.DocumentLine{})
	db.AutoMigrate(&models.Document{})

	var user models.User
	if err := db.Where("username = ?", "admin_axyz").First(&user).Error; err != nil {

		password, err := utils.HashPassword("OcdcT0TLVZ")
		if err != nil {
			log.Fatal(err)
		}
		db.Create(&models.User{
			Username: "admin_axyz",
			Password: password,
		})
	}

}

func GetDB() *gorm.DB {
	return DBPg
}
