package db

import (
	"SPMA-APP/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB
var database *gorm.DB
var err error

func Init() {
	conf := config.GetConfig()
	ConfigurationString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err = sql.Open("mysql", ConfigurationString)
	if err != nil {
		panic("connection error")
	}

	err := db.Ping()

	if err != nil {
		fmt.Println(err)
		panic("DSN error")
	}
}

func DatabaseInit() {
	conf := config.GetConfig()

	dsn := "host=" + conf.DB_HOST + " user=" + conf.DB_USERNAME + " password=" + conf.DB_PASSWORD + " dbname=" + conf.DB_NAME + " port=" + conf.DB_PORT + " sslmode=disable TimeZone=Asia/Jakarta"

	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}

func CreateCon() *sql.DB {
	return db
}

func CreateConGorm() *gorm.DB {
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	return gormDB
}
