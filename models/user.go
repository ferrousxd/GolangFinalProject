package models

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