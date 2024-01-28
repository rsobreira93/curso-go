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

	product := Product{Name: "Notebook", Price: 10.0, CategoryID: 1}

	db.Create(&product)

	db.Create(&SerialNumber{Number: "123456789", ProductID: 1})

	var categories []Category

	// err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			println("- ", p.Name, "Serial Number: ", p.SerialNumber.Number)
		}
	}

}
