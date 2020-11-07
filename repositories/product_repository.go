package repositories

import (
	"GolangFinalProject/models"
	"database/sql"
)

type ProductRepository struct {
	connection *sql.DB
}

func (userRepo *UserRepository) InsertProduct(product models.Product) {

}

func (userRepo *UserRepository) DeleteProduct(productId int) {

}

func (userRepo *UserRepository) GetAllProducts() []*models.Product {
	return nil
}
//Test