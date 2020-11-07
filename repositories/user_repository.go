package repositories

import (
	"GolangFinalProject/models"
	"database/sql"
)

type UserRepository struct {
	connection *sql.DB
}

func (userRepo *UserRepository) InsertUser(user models.User) {

}

func (userRepo *UserRepository) GetUserByLogin(username string, password string) *models.User {
	return nil
}