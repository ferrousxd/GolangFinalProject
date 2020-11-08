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
	_, err :=  pr.Connection.Exec(`DELETE FROM products WHERE id = $1 `, productId)

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

func (pr *ProductRepository) GetProductById(productId int) *models.Product {
	productBuilder := models.ProductBuilder{}

	rows, err := pr.Connection.Query(`SELECT id, model, company, price FROM products WHERE id = $1`, productId)
  
	if err != nil {
		panic(err)
	}

	var product *models.Product

	if rows.Next() {
		var id 		int
		var model 	string
		var company string
		var price 	float32

		err := rows.Scan(&id, &model, &company, &price)

		if err != nil {
			fmt.Println(err)
		}

		product = productBuilder.
			SetId(id).
			SetModel(model).
			SetCompany(company).
			SetPrice(price).
			Build()
	}
  
	return product
}

func (pr *ProductRepository) GetProductsBySubscriberId(userId int) []*models.Product {
	rows, err := pr.Connection.Query("SELECT p.id, p.model, p.company, p.price FROM products p INNER JOIN subscriptions s on p.id = s.product_id INNER JOIN users u on u.id = s.user_id WHERE u.id = $1", userId)

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
			panic(err)
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