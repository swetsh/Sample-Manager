package config

import (
	"fmt"
	"log"
	"sample-manager/model"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbname := "samplemanagerdb"
	dbuser := "postgres"
	password := "0000"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbuser,
		dbname,
		password,
	)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	d.AutoMigrate(&model.SampleItem{}, &model.Segment{})

	if err != nil {
		log.Fatal(err)
	}

	db = d

	fmt.Println("Database connection successful...")
}

func GetDB() *gorm.DB {
	return db
}
