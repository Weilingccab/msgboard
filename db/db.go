package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type dbConfig struct {
	Addr     string
	Port     int
	Username string
	Name     string
	Password string
}

func getDbConfig() *dbConfig {
	log.Println("Starting server..")

	config := dbConfig{}
	file := "db/config.json"
	data, err := ioutil.ReadFile(file)
	json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return &config
}

func connectDB() *gorm.DB {
	var err error
	config := getDbConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Addr,
		config.Port,
		config.Username,
		config.Name,
		config.Password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return db
	}
	return db
}

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}
