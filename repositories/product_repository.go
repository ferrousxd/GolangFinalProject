package repositories

import (
	"GolangFinalProject/models"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	Connection *sql.DB
}

func (pr *ProductRepository) InsertProduct(product models.Product) {
	_, err := pr.Connection.Exec(`INSERT INTO products(model, company, price) VALUES ($1, $2, $3)`,
		product.GetModel(), product.GetCompany(), product.GetPrice())

	if err != nil {
		panic(err)
	}
}

func (pr *ProductRepository) DeleteProduct(productId int) {
	_, err :=  pr.Connection.Exec(`DELETE FROM products where id = $1 `, productId)

	if err != nil {
		panic(err)
	}
}

func (pr *ProductRepository) GetAllProducts() []*models.Product {
	rows, err := pr.Connection.Query(`SELECT id, model, company, price FROM products`)

	if err != nil {
		panic(err)
	}

	var products []*models.Product

	for rows.Next() {
		productBuilder := models.ProductBuilder{}

		var id int
		var model string
		var company string
		var price float32

		err := rows.Scan(&id, &model, &company, &price)

		if err != nil {
			fmt.Println(err)
			continue
		}

		product := productBuilder.
			SetId(id).
			SetModel(model).
			SetCompany(company).
			SetPrice(price).
			Build()

		products = append(products, product)
	}

	return products
}