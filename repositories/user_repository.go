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

	rows, err := ur.Connection.Query(`SELECT id, username, email, role, balance FROM users where username = $1 and password = $2`,
		username, password)

	if err != nil {
		panic(err)
	}

	var user *models.User

	if rows.Next(){
		var id int
		var username, email, role string
		var balance float32

		err := rows.Scan(&id, &username, &email, &role, &balance)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		user = userBuilder.
			SetId(id).
			SetUsername(username).
			SetEmail(email).
			SetRole(role).
			SetBalance(balance).
			Build()
	}

	return user
}

func (ur *UserRepository) GetUserById(userId int) *models.User {
	userBuilder := models.UserBuilder{}

	rows, err := ur.Connection.Query(`SELECT id, username, email, role, balance FROM users where id = $1`,
		userId)

	if err != nil {
		panic(err)
	}

	var user *models.User

	if rows.Next(){
		var id int
		var username, email, role string
		var balance float32

		err := rows.Scan(&id, &username, &email, &role, &balance)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		user = userBuilder.
			SetId(id).
			SetUsername(username).
			SetEmail(email).
			SetRole(role).
			SetBalance(balance).
			Build()
	}

	return user
}

func (ur *UserRepository) ChangeSubscriptionStatus(productId int, userId int, operation string) {
	if operation == "add" {
		_, err := ur.Connection.Exec("INSERT INTO subscriptions(product_id, user_id) VALUES ($1, $2)", productId, userId)

		if err != nil {
			panic(err)
		}
	} else if operation == "remove" {
		_, err := ur.Connection.Exec("DELETE FROM subscriptions WHERE product_id = $1 AND user_id = $2", productId, userId)

		if err != nil {
			panic(err)
		}
	}
}

func (ur *UserRepository) GetSubscribersByProductId(productId int) []*models.User {
	rows, err := ur.Connection.Query(`SELECT DISTINCT u.id, u.username, u.email FROM users u INNER JOIN subscriptions s on u.id = s.user_id INNER JOIN products p on p.id = s.product_id WHERE s.product_id = $1`, productId)

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

func (ur *UserRepository) AddMoneyToBalance(u *models.User, plusAmount float32) {
	if plusAmount > 0 {
		_, err := ur.Connection.Exec(`UPDATE users SET balance = $1 WHERE email = $2`, u.GetBalance() + plusAmount, u.GetEmail())

		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("ALO A GDE DEN'GI")
	}
}

func (ur *UserRepository) RemoveMoneyFromBalance(u *models.User, minusAmount float32) bool {
	if minusAmount > 0 && u.GetBalance() >= minusAmount {
		_, err := ur.Connection.Exec(`UPDATE users SET balance = $1 WHERE email = $2`, u.GetBalance() - minusAmount, u.GetEmail())

		if err != nil {
			panic(err)
		}

		return true
	} else {
		return false
	}
}