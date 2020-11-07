package models

import (
	"fmt"
	"net/smtp"
)

type observer interface {
	update(string)
}

type User struct {
	id                 int
	username           string
	email              string
	password           string
	role               string
	subscriptionStatus bool
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

func (b *UserBuilder) SetSubscriptionStatus(subscriptionStatus bool) *UserBuilder {
	b.actions = append(b.actions, func(u *User) {
		u.subscriptionStatus = subscriptionStatus
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

func (u *User) GetSubscriptionStatus() bool {
	return u.subscriptionStatus
}

func (u *User) GetBalance() float32 {
	return u.balance
}

func (u *User) update(model string) {
	from := "superusergoproject@gmail.com"
	password := "imyaMoyeiSobaki"

	to := []string {
		u.GetEmail(),
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("To: <" + to[0] + ">\r\n" +
		"From: Chechnya Bank Admin\n" +
		"Subject: New product has arrived!\n" +
		"\n" +
		"Hello " + to[0] + "! We hope you are doing great. " + "Now, " + model + " is available in our shop!\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	for i := 0; i < 100; i++ {
		err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, to, message)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Message was sent successfully")
	}
}