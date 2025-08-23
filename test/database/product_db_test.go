package database

import (
	"fmt"
	"math/rand"
	"projeto-modelo/internal/entity"
	"projeto-modelo/internal/infra/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductDBTestSuite struct {
	suite.Suite
	db        *gorm.DB
	productDB *database.Product
}

func (suite *ProductDBTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	suite.Require().NoError(err, "failed to connect to database")

	err = db.AutoMigrate(&entity.Product{})
	suite.Require().NoError(err, "failed to migrate database")

	suite.db = db
	suite.productDB = database.NewProduct(db)
}

func (suite *ProductDBTestSuite) TearDownTest() {
	if suite.db != nil {
		sqlDB, err := suite.db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}

func (suite *ProductDBTestSuite) createTestProduct(name string, price int) *entity.Product {
	product, err := entity.NewProduct(name, price)
	suite.Require().NoError(err, "failed to create product entity")

	err = suite.db.Create(product).Error
	suite.Require().NoError(err, "failed to save product to database")

	return product
}

func (suite *ProductDBTestSuite) createMultipleProducts(count int) []*entity.Product {
	products := make([]*entity.Product, count)
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("Product %d", i+1)
		price := rand.Intn(100) + 1
		products[i] = suite.createTestProduct(name, price)
	}
	return products
}

func TestProductDBTestSuite(t *testing.T) {
	suite.Run(t, new(ProductDBTestSuite))
}

func (suite *ProductDBTestSuite) TestCreateNewProduct() {
	product, err := entity.NewProduct("Product 1", 100)
	assert.NoError(suite.T(), err)

	err = suite.productDB.Create(product)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), product.ID)
	assert.Equal(suite.T(), "Product 1", product.Name)
	assert.Equal(suite.T(), 100, product.Price)
}

func (suite *ProductDBTestSuite) TestFindAllProducts() {
	suite.createMultipleProducts(23)

	products, err := suite.productDB.FindAll(1, 10, "asc")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), products, 10)
	assert.Equal(suite.T(), "Product 1", products[0].Name)
	assert.Equal(suite.T(), "Product 10", products[9].Name)

	products, err = suite.productDB.FindAll(2, 10, "asc")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), products, 10)
	assert.Equal(suite.T(), "Product 11", products[0].Name)
	assert.Equal(suite.T(), "Product 20", products[9].Name)

	products, err = suite.productDB.FindAll(3, 10, "asc")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), products, 3)
	assert.Equal(suite.T(), "Product 21", products[0].Name)
	assert.Equal(suite.T(), "Product 23", products[2].Name)
}

func (suite *ProductDBTestSuite) TestFindProductByID() {
	createdProduct := suite.createTestProduct("Product 1", 100)

	foundProduct, err := suite.productDB.FindByID(createdProduct.ID.String())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), createdProduct.Name, foundProduct.Name)
	assert.Equal(suite.T(), createdProduct.Price, foundProduct.Price)

	nonExistentProduct, err := suite.productDB.FindByID("non-existent-id")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), nonExistentProduct)
}

func (suite *ProductDBTestSuite) TestUpdateProduct() {
	product := suite.createTestProduct("Product 1", 100)

	product.Name = "Product 2"
	product.Price = 200

	err := suite.productDB.Update(product.ID.String(), product)
	assert.NoError(suite.T(), err)

	updatedProduct, err := suite.productDB.FindByID(product.ID.String())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Product 2", updatedProduct.Name)
	assert.Equal(suite.T(), 200, updatedProduct.Price)

	product.Name = "Product 3"
	product.Price = 300
	err = suite.productDB.Update(product.ID.String(), product)
	assert.NoError(suite.T(), err)

	updatedProduct, err = suite.productDB.FindByID(product.ID.String())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Product 3", updatedProduct.Name)
	assert.Equal(suite.T(), 300, updatedProduct.Price)
}

func (suite *ProductDBTestSuite) TestDeleteProduct() {
	product := suite.createTestProduct("Product 1", 100)

	err := suite.productDB.Delete(product.ID.String())
	assert.NoError(suite.T(), err)

	deletedProduct, err := suite.productDB.FindByID(product.ID.String())
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), deletedProduct)
}
