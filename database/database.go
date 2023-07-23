package database

import (
	"fmt"
	"log"
	"os"

	"github.com/efealtar/gofib/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations...")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	} // set global DB variable to db
}
