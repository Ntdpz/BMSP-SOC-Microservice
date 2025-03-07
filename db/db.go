package db

import (
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

func GetDB() *gorm.DB {
	return DBPg
}
