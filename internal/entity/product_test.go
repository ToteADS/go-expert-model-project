package entity

import (
	"projeto-modelo/pkg/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Nil(t, product.Validate())
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Product 1", 100)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}

func TestProductValidateWithInvalidID(t *testing.T) {
	product := Product{
		ID:        entity.ID{}, 
		Name:      "Product 1",
		Price:     100,
		CreatedAt: time.Now(),
	}
	assert.Equal(t, ErrIDIsRequired, product.Validate())
}
