package database

import (
	"projeto-modelo/internal/entity"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindAll(page, limit int, sort string) ([]entity.User, error) {
	var users []entity.User
	var offset int
	if page > 0 {
		offset = (page - 1) * limit
	}
	
	// Use a valid field for ordering since created_at doesn't exist
	orderBy := "id"
	if sort != "" {
		orderBy = sort
	}
	
	tx := u.DB.Limit(limit).Offset(offset).Order(orderBy).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	tx := u.DB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u *User) FindByID(id string) (*entity.User, error) {
	var user entity.User
	tx := u.DB.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u *User) Update(user *entity.User) error {
	return u.DB.Save(user).Error
}

func (u *User) Delete(id string) error {
	return u.DB.Where("id = ?", id).Delete(&entity.User{}).Error
}
