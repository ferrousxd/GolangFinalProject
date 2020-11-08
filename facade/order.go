package facade

import (
	"GolangFinalProject/models"
	"GolangFinalProject/repositories"
	"fmt"
	"strings"
)

type orderFacade struct {
	user 			*models.User
	balance 		float32
	notification 	*notification
	products 		[]*models.Product
	userRepo		repositories.UserRepository
}

func NewOrderFacade(u *models.User, ps []*models.Product, ur repositories.UserRepository) *orderFacade {
	fmt.Println("Starting to process order...")
	orderFacade := &orderFacade{
		user:         u,
		balance:      u.GetBalance(),
		notification: &notification{},
		products:     ps,
		userRepo:     ur,
	}

	if ps == nil {
		fmt.Println("You have no items!")
	} else {
		fmt.Println("Total price of order with discount: ", calculateTotalPrice(ps))
		orderFacade.PrintProduct()
	}

	return orderFacade
}

func (of *orderFacade) PrintProduct(){
	count := 1
	for _, p := range of.products {
		fmt.Println(count, " | ", p.GetId(), " | ", p.GetCompany(), " | " + p.GetModel(), " | ", p.GetPrice())
		count++
	}
}

func (of *orderFacade) RemoveFromOrder(countTable int) {
	s := of.products
	productListLength := len(s)
	for i, _ := range s {
		if countTable == i {
			s[productListLength-1], s[i] = s[i], s[productListLength-1]
		}
	}
	of.products = s[:productListLength-1]

	//s := *of.products
	//s = append(s[:id], s[id+1:]...)
	//*of.products = s
}

func calculateTotalPrice(ps []*models.Product) float32 {
	var totalPrice float32
	var totalDiscount float32

	for _, p := range ps {
		if p.GetPrice() > 0 && p.GetPrice() < 10 {
			totalDiscount += 30.0
		} else if p.GetPrice() > 10 && p.GetPrice() < 100 {
			totalDiscount += 20.0
		} else if p.GetPrice() > 100 {
			totalDiscount += 10.0
		} else {
			totalDiscount += 5
		}

		if strings.Contains(p.GetModel(),"S2") {
			totalDiscount += 30.0
		} else if strings.Contains(p.GetModel(),"N320") {
			totalDiscount += 20.0
		} else if strings.Contains(p.GetModel(),"iPhone") {
			totalDiscount += 10.0
		} else {
			totalDiscount += 5
		}

		if p.GetCompany() == "Xiaomi" {
			totalDiscount += 30.0
		} else if p.GetCompany() == "Samsung" {
			totalDiscount += 20.0
		} else if p.GetCompany() == "Apple" {
			totalDiscount += 10.0
		} else {
			totalDiscount += 5
		}
		totalPrice = totalPrice + p.GetPrice() * (100 - totalDiscount) / 100
		// Обнуляем discount
		totalDiscount = 0
	}

	return totalPrice
}

func (of *orderFacade) MakeOrder() {
	if of.userRepo.RemoveMoneyFromBalance(of.user, calculateTotalPrice(of.products)) == true {
		of.notification.sendSuccessOrderNotification()
		of.products = nil
	} else {
		fmt.Println("ALO A CHE TAM S DEN'GAMI")
	}
}

func (of *orderFacade) CancelOrder() {
	of.notification.sendCloseOrderNotification()
}

type notification struct {}

func (n *notification) sendSuccessOrderNotification() {
	fmt.Println("Order was made successfully!")
}

func (n *notification) sendCloseOrderNotification() {
	fmt.Println("Order was canceled successfully!")
}

//func newDiscountFacade(product models.Product) *discountFacade  {
//	productPrice := product.GetPrice()
//	productModel := product.GetModel()
//	productCompany := product.GetCompany()
//	fmt.Println("Starting to run discount...")
//	discountFacade := &discountFacade{
//		price: discountByPrice(productPrice),
//		model: discountByModel(productModel),
//		company: discountByCompany(productCompany),
//	}
//	fmt.Println("Discount is on ran")
//}