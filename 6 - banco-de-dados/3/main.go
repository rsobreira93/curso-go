package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Product struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category `gorm:"foreignKey:CategoryID"`
	gorm.Model
}

func main() {
	dsn := "host=localhost user=fon password=fon dbname=fonfon port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Eletronicos"}

	// db.Create(&category)

	// product := Product{Name: "Notebook", Price: 10.0, CategoryID: category.ID}

	// db.Create(&product)

	var products []Product

	db.Preload("Category").Find(&products)

	for _, p := range products {
		fmt.Println(p)
	}

}
