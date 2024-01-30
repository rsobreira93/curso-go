package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.5)

	assert.Nil(t, err)

	assert.NotNil(t, product)

	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Price)
	assert.NotEmpty(t, product.Created_at)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10.5)

	assert.Nil(t, product)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)

	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -10.5)

	assert.Nil(t, product)
	assert.Equal(t, ErrPriceInvalid, err)
}

func TestProductValidadte(t *testing.T) {
	p, err := NewProduct("Product 1", 10.5)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
