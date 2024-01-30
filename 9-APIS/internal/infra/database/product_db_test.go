package database

import (
	"fmt"
	"testing"

	"github.com/rsobreira93/curso-go/9-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.5)

	productDB := NewProduct(db)

	err = productDB.Create(product)

	assert.Nil(t, err)

	var productFound entity.Product

	err = db.First(&productFound, "id = $1", product.ID).Error

	assert.Nil(t, err)

	assert.NotEmpty(t, productFound.ID)

	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.5)

	productDB := NewProduct(db)

	err = productDB.Create(product)

	assert.Nil(t, err)

	productFound, err := productDB.FindById(product.ID.String())

	assert.Nil(t, err)

	assert.Equal(t, "Product 1", productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.5)

	productDB := NewProduct(db)

	err = productDB.Create(product)

	assert.Nil(t, err)

	product.Name = "Product 2"
	product.Price = 20.5

	err = productDB.Update(product)

	assert.Nil(t, err)

	var productFound entity.Product

	err = db.First(&productFound, "id = $1", product.ID).Error

	assert.Nil(t, err)

	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 2", productFound.Name)
	assert.Equal(t, 20.5, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.5)

	productDB := NewProduct(db)

	err = productDB.Create(product)

	assert.Nil(t, err)

	err = productDB.Delete(product.ID.String())

	assert.Nil(t, err)

	// err = db.First(product.ID, "id = $1", product.ID).Error
	_, err = productDB.FindById(product.ID.String())

	assert.Error(t, err)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i <= 25; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), float64(i))
		err = db.Create(product).Error
		assert.Nil(t, err)
	}

	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")

	assert.NoError(t, err)

	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 25", products[4].Name)

}
