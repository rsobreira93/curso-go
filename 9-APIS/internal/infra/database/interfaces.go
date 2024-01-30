package database

import "github.com/rsobreira93/curso-go/9-APIS/internal/entity"

type UserInterface interface {
	Create(User *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(Product *entity.Product) error
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(Product *entity.Product) error
	Delete(id string) error
}
