package main

import (
	"encoding/json"
	"net/http"

	"github.com/rsobreira93/curso-go/9-APIS/configs"
	"github.com/rsobreira93/curso-go/9-APIS/internal/dto"
	"github.com/rsobreira93/curso-go/9-APIS/internal/entity"
	"github.com/rsobreira93/curso-go/9-APIS/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	http.ListenAndServe(":8080", nil)

}

type ProductHandler struct {
	ProductDb database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDb: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductDTO

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, error := entity.NewProduct(product.Name, product.Price)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDb.Create(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
