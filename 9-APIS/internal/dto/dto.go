package dto

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
