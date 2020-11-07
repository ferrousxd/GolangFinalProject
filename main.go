package main

import (
	"GolangFinalProject/models"
	"GolangFinalProject/repositories"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	//var input string
	//fmt.Scanln(&input)
	//fmt.Println(input)

	db := repositories.GetSingletonDatabase()
	dbConn := db.GetConnection()

	productRepo := repositories.ProductRepository{ dbConn}

	productBuilder := models.ProductBuilder{}

	product := productBuilder.SetCompany("Samsung").SetPrice(42.1).SetModel("S2").Build()

	productRepo.InsertProduct(*product)
	productRepo.InsertProduct(*(
		productBuilder.
			SetCompany("Apple").
			SetModel("Iphone X").
			SetPrice(42.99).
			Build()))

	allProducts := productRepo.GetAllProducts()

	for _, p := range allProducts {
		fmt.Println(p.GetId(), p.GetModel(), p.GetPrice(), p.GetCompany())
	}

	productRepo.DeleteProduct(1)

	userRepo := repositories.UserRepository{dbConn}

	userBuilder := models.UserBuilder{}

	userRepo.InsertUser(*(
			userBuilder.
				SetUsername("madiyar").
				SetPassword("123456").
				SetEmail("madok@gmail.com").
				Build()))

	userRepo.InsertUser(*(
		userBuilder.
			SetUsername("chingiz").
			SetPassword("123456").
			SetEmail("chinga@gmail.com").
			Build()))

	userRepo.InsertUser(*(
		userBuilder.
			SetUsername("azatkali").
			SetPassword("123456").
			SetEmail("azaza@gmail.com").
			Build()))

	fmt.Println(userRepo.GetUserByLogin("madiyar", "123456").GetUsername(),
		userRepo.GetUserByLogin("madiyar", "123456").GetEmail(),
		userRepo.GetUserByLogin("madiyar", "123456").GetStatus())

	fmt.Println(userRepo.GetUserByLogin("chingiz", "123456").GetUsername(),
		userRepo.GetUserByLogin("chingiz", "123456").GetEmail(),
		userRepo.GetUserByLogin("chingiz", "123456").GetStatus())

	fmt.Println(userRepo.GetUserByLogin("azatkali", "123456").GetUsername(),
		userRepo.GetUserByLogin("azatkali", "123456").GetEmail(),
		userRepo.GetUserByLogin("azatkali", "123456").GetStatus())

}