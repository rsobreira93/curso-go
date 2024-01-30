package entity

import (
	"errors"
	"time"

	"github.com/rsobreira93/curso-go/9-APIS/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("iD is required")
	ErrInvalidID       = errors.New("invalid ID")
	ErrNameRequired    = errors.New("name is required")
	ErrPriceInvalid    = errors.New("price is invalid")
	ErrPriceIsRequired = errors.New("price is required")
)

type Product struct {
	ID         entity.ID `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Created_at time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:         entity.NewID(),
		Name:       name,
		Price:      price,
		Created_at: time.Now(),
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrPriceInvalid
	}

	return nil
}
