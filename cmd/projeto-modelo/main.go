package main

import (
	"encoding/json"
	"log"
	"net/http"
	"projeto-modelo/configs"
	"projeto-modelo/internal/dto"
	"projeto-modelo/internal/entity"
	"projeto-modelo/internal/infra/database"

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

    http.ListenAndServe(":8000", nil)
}

//Criando os handlers mas não ficará aqui, será criado em um arquivo separado.
type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product dto.CreateProductInput
    err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    //De forma geral o handler não deve ter acesso a entidade, normalmente isso deve ser feito pelo caso de uso.
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(p)
}