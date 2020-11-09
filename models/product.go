package models

// Observable
type subject interface {
	AddObserver(Observer observer)
	RemoveObserver(Observer observer)
	NotifyAllObservers()
}

// Component interface
type Decorator interface {
	GetId() int
	GetModel() string
	GetCompany() string
	GetPrice() float32
}
// Concrete Component and concrete Observable (2 в одном)
type Product struct {
	id      	 int
	model   	 string
	company 	 string
	price   	 float32
	observerList []observer
}

// Concrete Decorator - 128GB
type With128GB struct {
	Decorator Decorator
}

func (w128 *With128GB) GetId() int {
	return w128.Decorator.GetId()
}

func (w128 *With128GB) GetModel() string {
	return w128.Decorator.GetModel()
}

func (w128 *With128GB) GetCompany() string {
	return w128.Decorator.GetCompany()
}

func (w128 *With128GB) GetPrice() float32 {
	return w128.Decorator.GetPrice()*1.1
}

// Concrete Decorator - 256GB
type With256GB struct {
	Decorator Decorator
}

func (w256 *With256GB) GetId() int {
	return w256.Decorator.GetId()
}

func (w256 *With256GB) GetModel() string {
	return w256.Decorator.GetModel()
}

func (w256 *With256GB) GetCompany() string {
	return w256.Decorator.GetCompany()
}

func (w256 *With256GB) GetPrice() float32 {
	return w256.Decorator.GetPrice()*1.25
}

// Concrete Decorator - Phone Case
type WithCase struct {
	Decorator Decorator
}

func (c *WithCase) GetId() int {
	return c.Decorator.GetId()
}

func (c *WithCase) GetModel() string {
	return c.Decorator.GetModel()
}

func (c *WithCase) GetCompany() string {
	return c.Decorator.GetCompany()
}

func (c *WithCase) GetPrice() float32 {
	return c.Decorator.GetPrice()*1.05
}

// Builder with Fluent Interfaces
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

// Getters for Concrete Component
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

// Observable methods
func (p *Product) AddObserver(o observer) {
	p.observerList = append(p.observerList, o)
}

func (p *Product) RemoveObserver(o observer) {
	p.observerList = removeFromSlice(p.observerList, o)
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