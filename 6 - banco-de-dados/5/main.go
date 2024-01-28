package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product
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

	category := Category{Name: "Eletronicos"}

	db.Create(&category)

	product := Product{Name: "Notebook", Price: 10.0, CategoryID: category.ID}

	db.Create(&product)

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name)
		for _, p := range c.Products {
			println(p.Name)
		}
	}

}
