package models

import (
	"fmt"
	"net/smtp"
)

type observer interface {
	Notify(string)
	GetId() int
}

type User struct {
	id                 int
	username           string
	email              string
	password           string
	role               string
	balance			   float32
}

type userMod func(*User)

type UserBuilder struct {
	actions []userMod
}

func (b *UserBuilder) SetId(id int) *UserBuilder {
	b.actions = append(b.actions, func(u *User) {
		u.id = id
	})
	return b
}

func (b *UserBuilder) SetUsername(username string) *UserBuilder {
	b.actions = append(b.actions, func(u *User) {
		u.username = username
	})
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.actions = append(b.actions, func(u *User) {
		u.email = email
	})
	return b
}

func (b *UserBuilder) SetPassword(password string) *UserBuilder {
	b.actions = append(b.actions, func(u *User) {
		u.password = password
	})
	return b
}

func (b *UserBuilder) SetRole(role string) *UserBuilder {
	b.actions = append(b.actions, func(u *User) {
		u.role = role
	})
	return b
}

func (b *UserBuilder) SetBalance(balance float32) *UserBuilder {
	b.actions = append(b.actions, func(u *User) {
		u.balance = balance
	})
	return b
}

func (b *UserBuilder) Build() *User {
	user := &User{}

	for _, i := range b.actions {
		i(user)
	}

	return user
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetRole() string {
	return u.role
}

func (u *User) GetBalance() float32 {
	return u.balance
}

func (u *User) Notify(model string) {
	from := "superusergoproject@gmail.com"
	imyaMoyeiSobaki := "imyaMoyeiSobaki"

	to := []string {
		u.GetEmail(),
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Hello, " + u.GetEmail() + "! Check out the brand new " + model + " in our shop!")

	auth := smtp.PlainAuth("", from, imyaMoyeiSobaki, smtpHost)

	for i := 0; i < 1; i++ {
		err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Сообщение было успешно отправлено")
	}
}