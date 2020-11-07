package main

import (
	"GolangFinalProject/models"
	"GolangFinalProject/repositories"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)


	db := repositories.GetSingletonDatabase()
	dbConn := db.GetConnection()

	rows, err := dbConn.Query("select * from Products")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	fmt.Println(rows)

	var products []*models.Product

	for rows.Next() {
		productBuilder := models.ProductBuilder{}

		var id 		int
		var model 	string
		var company string
		var price 	float32

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

	for _, p := range products {
		fmt.Println(p.GetId(), p.GetModel(), p.GetPrice(), p.GetCompany())
	}
}