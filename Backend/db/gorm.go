package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connString := os.Getenv("SERVER_DB")
	database, err := gorm.Open(postgres.Open(fmt.Sprintf(connString, os.Getenv("USER_DB"), os.Getenv("PASS_DB"), os.Getenv("PORT_DB"))), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: ", err)
	}
	DB = database
}
