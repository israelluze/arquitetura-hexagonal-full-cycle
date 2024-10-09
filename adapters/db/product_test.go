package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/israelluze/go-hexagonal/adapters/db"
	"github.com/israelluze/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func Setup() {
	var err error
	Db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	createProductTable := `
	CREATE TABLE products (
		"id" string ,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(createProductTable)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insertProduct := `
	INSERT INTO products VALUES ("1", "Product 1", 0, "disabled");
	`
	stmt, err := db.Prepare(insertProduct)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	Setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("1")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	Setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Status, productResult.GetStatus())
}
