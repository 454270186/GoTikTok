package dal

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var (
	dbHost string
	dbName string
	dbUser string
	dbPsw  string
	dbPort string

	dsn string
)

func init() {
	dbEnv, err := godotenv.Read("../.env")
	if err != nil {
		panic("fail to read db env: " + err.Error())
	}

	dbHost = dbEnv["DB_HOST"]
	dbName = dbEnv["DB_NAME"]
	dbUser = dbEnv["DB_USER"]
	dbPsw = dbEnv["DB_PSW"]
	dbPort = dbEnv["DB_PORT"]

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPsw, dbHost, dbPort, dbName)
}

func newDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	DB = initDB()

	if DB == nil {
		log.Fatal("BD is nil")
	}

	return DB
}

func initDB() *gorm.DB {
	//dsn := "xiaofei:2021110003@tcp(127.0.0.1:8091)/Tiktok?parseTime=true"
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
