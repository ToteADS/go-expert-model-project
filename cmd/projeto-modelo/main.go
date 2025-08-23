package main

import (
	"log"
	"net/http"
	"projeto-modelo/configs"
	"projeto-modelo/internal/entity"
	"projeto-modelo/internal/infra/database"
	"projeto-modelo/internal/infra/webserver/handlers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDb := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDb)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}
