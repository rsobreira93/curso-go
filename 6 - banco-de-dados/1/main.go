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
		port     = 5432
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

	p, err := findProductById(database, product.ID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %v, possui o valor de: %.2f", p.Name, p.Price)

	products, err := finProducts(database)

	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Product: %v, possui o valor de: %.2f \n", p.Name, p.Price)
	}

	err = deleteProduct(database, p.ID)

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

func updateProduct(database *sql.DB, product *Product) error {
	statement, err := database.Prepare("update products set name = $1, price = $2 where id = $3")

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

	var product Product

	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func finProducts(database *sql.DB) ([]Product, error) {
	rows, err := database.Query("select id, name, price from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product

		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func deleteProduct(database *sql.DB, id string) error {
	statement, err := database.Prepare("delete from products where id = $1")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
