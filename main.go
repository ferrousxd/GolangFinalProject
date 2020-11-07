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

	productRepo := repositories.ProductRepository{ dbConn}

	productBuilder := models.ProductBuilder{}

	product := productBuilder.SetCompany("Samsung").SetPrice(42.1).SetModel("S2").Build()

	productRepo.InsertProduct(*product)
	allProducts := productRepo.GetAllProducts()

	for _, p := range allProducts {
		fmt.Println(p.GetId(), p.GetModel(), p.GetPrice(), p.GetCompany())
	}

	productRepo.DeleteProduct(1)

}