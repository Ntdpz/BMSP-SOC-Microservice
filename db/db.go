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

	// Open connection
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // ใช้ logger สำหรับดูคำสั่ง SQL

		// disable logger
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		panic(err)
	}

	// Set connection pool options
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get DB object:", err)
		panic(err)
	}

	// Set MaxIdleConns: จำนวน connection ที่ไม่ได้ใช้งาน
	sqlDB.SetMaxIdleConns(10)

	// Set MaxOpenConns: จำนวน connection สูงสุดที่สามารถเปิดได้
	sqlDB.SetMaxOpenConns(25)

	// Set ConnMaxLifetime: เวลาเชื่อมต่อที่ใช้ได้สูงสุด
	sqlDB.SetConnMaxLifetime(60)

	fmt.Println("Successfully connected to the database!")

	DBPg = DB

	return DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.DocumentLine{})
	db.AutoMigrate(&models.Document{})

	// find user admin
	var user models.User
	if err := db.Where("username = ?", "admin_axyz").First(&user).Error; err != nil {
		// create user admin
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
