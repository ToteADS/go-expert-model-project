package database

import "projeto-modelo/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindAll(page, limit int, sort string) ([]entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	Update(id string, user *entity.User) error
	Delete(id string) error
}


type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(id string, product *entity.Product) error
	Delete(id string) error
}