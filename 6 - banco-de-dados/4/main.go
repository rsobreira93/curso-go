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
	ID           int `gorm:"primary_key"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category `gorm:"foreignKey:CategoryID"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int
}

func main() {
	dsn := "host=localhost user=fon password=fon dbname=fonfon port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Eletronicos"}

	db.Create(&category)

	product := Product{Name: "Notebook", Price: 10.0, CategoryID: category.ID}

	db.Create(&product)

	db.Create(&SerialNumber{Number: "123456789", ProductID: 3})

	var products []Product

	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, p := range products {
		fmt.Println(p)
	}

}
