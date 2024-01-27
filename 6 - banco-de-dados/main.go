package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

	const (
		host     = "localhost"
		port     = 5444
		user     = "fon"
		password = "fon"
		dbname   = "fonfon"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer database.Close()

	product := NewProduct("Mouse", 199.90)

	err = insertProduct(database, product)

	if err != nil {
		panic(err)
	}

	fmt.Println("Product: " + product.Name + " inserted successfully!")

	product.Price = 299.90

	err = updateProduct(database, product)

	if err != nil {
		panic(err)
	}
}

func insertProduct(database *sql.DB, product *Product) error {
	statement, err := database.Prepare("insert into products(id, name, price) values($1, $2, $3)")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func updateProduct(dabatse *sql.DB, product *Product) error {
	statement, err := dabatse.Prepare("update products set name = $1, price = $2 where id = $3")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(product.Name, product.Price, product.ID)

	if err != nil {
		return err
	}

	return nil

}

func findProductById(database *sql.DB, id string) (*Product, error) {
	statement, err := database.Prepare("select id, name, price from products where id = $1")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRow(id)

	var productID string
	var name string
	var price float64

	err = row.Scan(&productID, &name, &price)

	product := &Product{
		ID:    productID,
		Name:  name,
		Price: price,
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}
