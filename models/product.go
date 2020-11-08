package models

import "GolangFinalProject/repositories"

type subject interface {
	AddObserver(Observer observer, userRepo repositories.UserRepository)
	RemoveObserver(Observer observer, userRepo repositories.UserRepository)
	NotifyAllObservers()
}

type Product struct {
	id      	 int
	model   	 string
	company 	 string
	price   	 float32
	observerList []observer
}

type productMod func(*Product)

type ProductBuilder struct {
	actions []productMod
}

func (b *ProductBuilder) SetId(id int) *ProductBuilder {
	b.actions = append(b.actions, func(p *Product) {
		p.id = id
	})
	return b
}

func (b *ProductBuilder) SetModel(model string) *ProductBuilder {
	b.actions = append(b.actions, func(p *Product) {
		p.model = model
	})
	return b
}

func (b *ProductBuilder) SetCompany(company string) *ProductBuilder {
	b.actions = append(b.actions, func(p *Product) {
		p.company = company
	})
	return b
}

func (b *ProductBuilder) SetPrice(price float32) *ProductBuilder {
	b.actions = append(b.actions, func(p *Product) {
		p.price = price
	})
	return b
}

func (b *ProductBuilder) Build() *Product {
	product := &Product{}

	for _, i := range b.actions {
		i(product)
	}

	return product
}

func (p *Product) GetId() int {
	return p.id
}

func (p *Product) GetModel() string {
	return p.model
}

func (p *Product) GetCompany() string {
	return p.company
}

func (p *Product) GetPrice() float32 {
	return p.price
}

func (p *Product) AddObserver(o observer, userRepo repositories.UserRepository) {
	p.observerList = append(p.observerList, o)
	userRepo.ChangeSubscriptionStatus(o.GetId(), "add")
}

func (p *Product) RemoveObserver(o observer, userRepo repositories.UserRepository) {
	p.observerList = removeFromSlice(p.observerList, o)
	userRepo.ChangeSubscriptionStatus(o.GetId(), "remove")
}

func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetId() == observer.GetId() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (p *Product) NotifyAllObservers() {
	for _, observer := range p.observerList {
		observer.Notify(p.model)
	}
}