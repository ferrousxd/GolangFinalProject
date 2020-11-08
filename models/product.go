package models

type subject interface {
	AddObserver(Observer observer)
	RemoveObserver(Observer observer)
	NotifyAllObservers()
	//Decorator
	GetPrice() float32
}

type Product struct {
	id      	 int
	model   	 string
	company 	 string
	price   	 float32
	observerList []observer
}

//Decorator
type With128GB struct {
	product Product
}

func (w128 *With128GB) GetPrice() float32 {
	return w128.product.GetPrice()*1.1
}

type With256GB struct {
	product Product
}

func (w256 *With256GB) GetPrice() float32 {
	return w256.product.GetPrice()*1.25
}

type Case struct {
	product Product
}

func (c *Case) GetPrice() float32 {
	return c.product.GetPrice()*1.05
}
//Decorator

//Builder Fluid
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
//Builder Fluid

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

//Observer
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
//Observer