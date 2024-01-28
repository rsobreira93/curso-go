package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "host=localhost user=postgres password=docker dbname=goexpert port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	product := Product{Name: "Product 1", Price: 10.0}

	db.Create(&product)

	// products := []Product{
	// 	{Name: "Product 1", Price: 10.0},
	// 	{Name: "Product 2", Price: 20.0},
	// 	{Name: "Product 3", Price: 30.0},
	// 	{Name: "Product 4", Price: 40.0},
	// }

	// db.Create(&products)

	// var product Product
	// db.First(&product, "price = ?", "20")

	// fmt.Printf("Product: %v\n", product)

	// var product []Product

	// db.Find(&product)

	// for _, p := range product {
	// 	fmt.Printf("Product: %v\n", p)
	// }

	// var products []Product

	// db.Limit(2).Offset(2).Find(&products)

	// for _, p := range products {
	// 	fmt.Printf("Product: %v\n", p)
	// }

	// var products []Product

	// db.Where("price > ?", 20).Find(&products)

	// for _, p := range products {
	// 	fmt.Printf("Product: %v\n", p)
	// }

	// var products []Product

	// db.Where("name ILIKE ?", "%product%").Find(&products)

	// for _, p := range products {
	// 	fmt.Printf("Product: %v\n", p)
	// }

	// var product Product
	// db.First(&product, "price = ?", "20")
	// fmt.Printf("Product: %v\n", product)

	// product.Name = "Product 2 Updated"
	// db.Save(&product)

	// db.First(&product, "price = ?", "20")
	// fmt.Printf("Product: %v\n", product)
	// db.Delete(&product)
}
