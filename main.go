package main

import (
	"GolangFinalProject/facade"
	"GolangFinalProject/repositories"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db := repositories.GetSingletonDatabase()
	dbConn := db.GetConnection()

	fmt.Println(dbConn)

	userRepo := repositories.UserRepository{Connection: dbConn}
	productRepo := repositories.ProductRepository{Connection: dbConn}

	//var username string
	//var password string
	//
	//fmt.Scan(&username)
	//fmt.Scan(&password)
	//
	//fmt.Println(username, password)
	//
	//user := userRepo.GetUserByLogin(username, password)
	//
	//if user == nil {
	//	panic("Wrong Credentials!!!")
	//}
	//
	//products := productRepo.GetAllProducts()
	//
	//for _, product := range products {
	//	fmt.Println(product.GetId(), " | ", product.GetCompany(), " | " + product.GetModel(), " | ", product.GetPrice())
	//}
	//
	//var productId int
	//fmt.Scan(&productId)
	//
	//users := userRepo.GetSubscribersByProductId(productId)
	//
	//for _, user := range users {
	//	fmt.Println(user.GetId(), " | ", user.GetUsername(), " | " + user.GetEmail())
	//}


	/*
		1.Add new Product
		2.Notify subscribers about the products

		2

		Choose the product, that you want to notify users about:
			1.
			2.
			3.

		1


	user1 := userRepo.GetUserByLogin("Chinga", "123")
	user2 := userRepo.GetUserByLogin("Aza", "456")
	user3 := userRepo.GetUserByLogin("Madik", "789")

	product1 := productRepo.GetProductById(1)
	product2 := productRepo.GetProductById(2)
	product3 := productRepo.GetProductById(3)

	product1.AddObserver(user1)
	product2.AddObserver(user2)
	product3.AddObserver(user3)

	product1.NotifyAllObservers()
	fmt.Println("1")
	product2.NotifyAllObservers()
	fmt.Println("2")
	product3.NotifyAllObservers()
	fmt.Println("3")

	 */

	username := "Chinga"
	password := "123"

	user1 := userRepo.GetUserByLogin(username, password)

	products := productRepo.GetAllProducts()

	for _, product := range products {
		fmt.Println(product.GetId())
	}

	//var products []*models.Product



	//userRepo.AddMoneyToBalance(user1, 1000)
	//Перезаписываем юзера после пополнения баланса
	user1 = userRepo.GetUserByLogin(username, password)

	newFacade := facade.NewOrderFacade(
		user1,
		products,
		userRepo,
	)

	var itemToRemove int

	fmt.Scan(&itemToRemove)
	newFacade.RemoveFromOrder(itemToRemove)

	newFacade.PrintProduct()

	newFacade.MakeOrder()

	// Which product you want to add to order(ID)
	// 1
	// productStandard := productRepo.GetProductById(1)
	// This product.GetModel has 64 GB of memory. Do you want to change the amount of storage?
	// Yes / No
	// Yes
	// 1. 128 GB
	// 2. 256 GB
	// If 128 / 256 -> productWithAdditional := With256{product: productStandard}

}