package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Status:", err)
	}

	fmt.Println("DB initialized!")

	return DB
}
