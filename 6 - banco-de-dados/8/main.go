package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	dsn := "host=localhost user=postgres password=docker dbname=goexpert port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Eletronicos"}

	// db.Create(&category)

	tx := db.Begin()

	var c Category

	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error

	if err != nil {
		panic(err)
	}

	c.Name = "EletronicoS"
	tx.Debug().Save(&c)
	tx.Commit()
}
