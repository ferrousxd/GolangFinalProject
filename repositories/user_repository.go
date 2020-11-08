package repositories

import (
	"GolangFinalProject/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	Connection *sql.DB
}

func (ur *UserRepository) InsertUser(user models.User) {
	_, err := ur.Connection.Exec(`INSERT INTO users(username, email, password) VALUES ($1, $2, $3)`,
		user.GetUsername(), user.GetEmail(), user.GetPassword())

	if err != nil {
		panic(err)
	}
}

func (ur *UserRepository) GetUserByLogin(username string, password string) *models.User {
	userBuilder := models.UserBuilder{}

	rows, err := ur.Connection.Query(`SELECT username, email, status FROM users where username = $1 and password = $2`,
		username, password)

	if err != nil {
		panic(err)
	}

	var user *models.User

	if rows.Next(){
		var username, email, status string

		err := rows.Scan(&username, &email, &status)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		user = userBuilder.
			SetUsername(username).
			SetEmail(email).
			SetStatus(status).
			Build()
	}

	return user
}

func (ur *UserRepository) ChangeSubscriptionStatus(userId int, operation string) {
	if operation == "add" {
		_, err := ur.Connection.Exec("UPDATE users SET subscription_status = true WHERE id = $1", userId)

		if err != nil {
			panic(err)
		}
	} else if operation == "remove" {
		_, err := ur.Connection.Exec("UPDATE users SET subscription_status = false WHERE id = $1", userId)

		if err != nil {
			panic(err)
		}
	}
}

func (ur *UserRepository) GetSubscribers() []*models.User {
	rows, err := ur.Connection.Query(`SELECT id, username, email FROM users WHERE subscription_status = true`)

	if err != nil {
		panic(err)
	}

	var users []*models.User

	for rows.Next() {
		userBuilder := models.UserBuilder{}

		var id int
		var username string
		var email string

		err := rows.Scan(&id, &username, &email)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		user := userBuilder.
			SetId(id).
			SetUsername(username).
			SetEmail(email).
			Build()

		users = append(users, user)
	}

	return users
}