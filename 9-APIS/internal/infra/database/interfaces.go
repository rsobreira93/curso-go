package database

import "github.com/rsobreira93/curso-go/9-APIS/internal/entity"

type UserInterface interface {
	Create(User *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
