package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/sousapedro11/fc-arquitetura-hexagonal/adapters/db"
	"github.com/sousapedro11/fc-arquitetura-hexagonal/application"
	"github.com/stretchr/testify/assert"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	createTable := `create table products (
		id string primary key,
		name string not null,
		price float,
		status string not null
	)`
	stmt, err := db.Prepare(createTable)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func createProduct(db *sql.DB) {
	insertProduct := `insert into products(id, name, price, status) values(?, ?, ?, ?)`
	stmt, err := db.Prepare(insertProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("1", "Product 1", 10.0, "enabled")
	if err != nil {
		log.Fatal(err)
	}
}

func TestProductDbGet(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	t.Run("Success - Get a product", func(t *testing.T) {
		result, err := productDb.Get("1")
		assert.Nil(t, err)
		assert.Equal(t, "Product 1", result.GetName())
		assert.Equal(t, 10.0, result.GetPrice())
		assert.Equal(t, "enabled", result.GetStatus())
	})

	t.Run("Error - Get a product that does not exist", func(t *testing.T) {
		result, err := productDb.Get("2")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestProductDbSave(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	t.Run("Success - Create a product", func(t *testing.T) {
		product := &application.Product{
			Id:     "2",
			Name:   "Product 2",
			Price:  20.0,
			Status: "enabled",
		}
		result, err := productDb.Save(product)
		assert.Nil(t, err)
		assert.Equal(t, "Product 2", result.GetName())
		assert.Equal(t, 20.0, result.GetPrice())
		assert.Equal(t, "enabled", result.GetStatus())
	})

	t.Run("Success - Update a product", func(t *testing.T) {
		product := &application.Product{
			Id:     "1",
			Name:   "Product 1",
			Price:  10.0,
			Status: "disabled",
		}
		result, err := productDb.Save(product)
		assert.Nil(t, err)
		assert.Equal(t, "Product 1", result.GetName())
		assert.Equal(t, 10.0, result.GetPrice())
		assert.Equal(t, "disabled", result.GetStatus())
	})
}
