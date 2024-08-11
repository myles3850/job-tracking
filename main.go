package main

import (
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Posting struct {
	ID          int
	JobTitle    string
	CompanyName string
	Link        *string
	Status      int
	Notes       *[]Notes
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Notes struct {
	ID          int
	PostingID int
	Note        string
	CreatedDate time.Time
	UpdatedDate time.Time
}

var envVars map[string]string

func init() {
	v, err := godotenv.Read()
	if err != nil {
		panic("cannot read env file")
	}
	envVars = v
}

func main() {
	db := NewDBConnection()
	db.AutoMigrate(&Posting{}, &Notes{})
}

func NewDBConnection() *gorm.DB {
	portStr := envVars["DB_PORT"]
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting DB_PORT to int: %v", err)
	}

	dsn := "host=" + envVars["DB_HOST"] + " user=" + envVars["DB_USER"] + " password=" + envVars["DB_PASS"] + " dbname=" + envVars["DB_NAME"] + " port=" + strconv.Itoa(portInt) + " sslmode=disable" + " TimeZone=" + "Europe/London"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err.Error())
	}
	return db

}
