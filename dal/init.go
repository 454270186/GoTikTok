package dal

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDB() *gorm.DB {
	DB := initDB()

	if DB == nil {
		log.Fatal("BD is nil")
	}

	return DB
}

func initDB() *gorm.DB {
	dsn := "xiaofei:2021110003@tcp(127.0.0.1:8091)/Tiktok?parseTime=true"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	log.Println("--------------------")
	log.Println("Mysql successfully init")
	log.Println("--------------------")

	return DB
}