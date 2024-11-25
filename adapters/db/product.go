package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sousapedro11/fc-arquitetura-hexagonal/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product application.Product
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	err := p.db.QueryRow("select count(id) from products where id = ?", product.GetId()).Scan(&rows)
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		err = p.create(product)
	} else {
		err = p.update(product)
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) create(product application.ProductInterface) error {
	stmt, err := p.db.Prepare("insert into products(id, name, price, status) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductDb) update(product application.ProductInterface) error {
	stmt, err := p.db.Prepare("update products set name = ?, price = ?, status = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {
		return err
	}

	return nil
}
