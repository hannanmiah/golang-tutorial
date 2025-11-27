package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/hannanmiah/golang-tutorial/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully!")
}