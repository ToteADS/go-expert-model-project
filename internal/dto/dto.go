package dto

import "projeto-modelo/pkg/entity"

type CreateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateProductOutput struct {
	ID    entity.ID `json:"id"`
	Name  string    `json:"name"`
	Price int       `json:"price"`
}
