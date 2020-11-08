package main

import (
	"GolangFinalProject/application"
	"GolangFinalProject/repositories"
	_ "github.com/lib/pq"
)

func main() {
	db := repositories.GetSingletonDatabase()
	dbConn := db.GetConnection()

	userRepo := repositories.UserRepository{Connection: dbConn}
	productRepo := repositories.ProductRepository{Connection: dbConn}

	newApplication := application.NewApplication(productRepo, userRepo)
	newApplication.Start()

}