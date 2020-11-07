package models

import (
	"fmt"
	"net/smtp"
)

type observer interface {
	update(string)
}

type User struct {
	id 		 int
	username string
	email 	 string
	password string
	status	 string
}

type userMod func(*User)

type UserBuilder struct {
	actions []userMod
}

func (b *UserBuilder) SetId(id int) *UserBuilder {
	b.actions = append(b.actions, func(p *User) {
		p.id = id
	})
	return b
}

func (b *UserBuilder) SetUsername(username string) *UserBuilder {
	b.actions = append(b.actions, func(p *User) {
		p.username = username
	})
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.actions = append(b.actions, func(p *User) {
		p.email = email
	})
	return b
}

func (b *UserBuilder) SetPassword(password string) *UserBuilder {
	b.actions = append(b.actions, func(p *User) {
		p.password = password
	})
	return b
}

func (b *UserBuilder) SetStatus(status string) *UserBuilder {
	b.actions = append(b.actions, func(p *User) {
		p.status = status
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

func (u *User) GetStatus() string {
	return u.status
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