package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rsobreira93/curso-go/9-APIS/configs"
	"github.com/rsobreira93/curso-go/9-APIS/internal/entity"
	"github.com/rsobreira93/curso-go/9-APIS/internal/infra/database"
	"github.com/rsobreira93/curso-go/9-APIS/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDb := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(productDb)

	userDb := database.NewUser(db)
	UserHandler := handlers.NewUserHandler(userDb, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", ProductHandler.CreateProduct)
	r.Get("/products", ProductHandler.GetAllProducts)
	r.Get("/products/{id}", ProductHandler.GetProduct)
	r.Put("/products/{id}", ProductHandler.UpdateProduct)
	r.Delete("/products/{id}", ProductHandler.DeleteProduct)

	r.Post("/users", UserHandler.CreateUser)
	r.Post("/users/login", UserHandler.GetJWT)

	http.ListenAndServe(":8080", r)

}
