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
	Products 		[]models.Decorator
	userRepo		repositories.UserRepository
}

func NewOrderFacade(u *models.User, ps []models.Decorator, ur repositories.UserRepository) *orderFacade {
	orderFacade := &orderFacade{
		user:         u,
		balance:      u.GetBalance(),
		notification: &notification{},
		Products:     ps,
		userRepo:     ur,
	}

	return orderFacade
}

func (of *orderFacade) PrintProduct() {
	rowNumber := 1
	if of.Products != nil {
		fmt.Println("Total price of products with discount: ", calculateTotalPrice(of.Products))
		for _, p := range of.Products {
			fmt.Println(rowNumber, " | ", p.GetId(), " | ", p.GetCompany(), " | " + p.GetModel(), " | ", p.GetPrice())
			rowNumber++
		}
	} else if of.Products == nil {
		fmt.Println("You have no items!")
	}
}

func (of *orderFacade) RemoveFromOrder(rowNumber int) {
	s := of.Products
	productListLength := len(s)
	for i, _ := range s {
		if rowNumber - 1 == i {
			s[productListLength-1], s[i] = s[i], s[productListLength-1]
		}
	}
	of.Products = s[:productListLength-1]
}

// calculateTotalPrice() receives slice of products and calculates total with discounts, which depend on manufacturer, model, and etc.
func calculateTotalPrice(ps []models.Decorator) float32 {
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

// If user has enough money - the order will be successfully executed - otherwise it will print out warning message
func (of *orderFacade) MakeOrder() {
	if of.userRepo.RemoveMoneyFromBalance(of.user, calculateTotalPrice(of.Products)) == true {
		fmt.Println("Starting to process order...")
		of.notification.sendSuccessOrderNotification()
		of.Products = nil
	} else {
		fmt.Println("ALO A CHE TAM S DEN'GAMI")
	}
}

func (of *orderFacade) CancelOrder() {
	of.Products = nil
	of.notification.sendCloseOrderNotification()
}

type notification struct {}

func (n *notification) sendSuccessOrderNotification() {
	fmt.Println("Order was made successfully!")
}

func (n *notification) sendCloseOrderNotification() {
	fmt.Println("Order was canceled successfully!")
}