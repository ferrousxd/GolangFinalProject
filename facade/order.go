package facade

import (
	"GolangFinalProject/models"
	"GolangFinalProject/repositories"
	"fmt"
)

type orderFacade struct {
	user 			*models.User
	balance 		float32
	notification 	*notification
	products 		*[]models.Product
	userRepo		repositories.UserRepository
}

func newOrderFacade(u *models.User, ps *[]models.Product, ur repositories.UserRepository) *orderFacade {
	fmt.Println("Starting to process order...")
	orderFacade := &orderFacade{
		user: u,
		balance: u.GetBalance(),
		notification: &notification{},
		products: ps,
		userRepo: ur,
	}
	fmt.Println("Total price of order with discount: ", calculateTotalPrice(ps))
	orderFacade.printProduct()
	return orderFacade
}

func (of *orderFacade) printProduct(){
	count := 1
	for _, p := range *of.products{
		fmt.Println(count, " | ", p.GetId(), " | ", p.GetCompany(), " | " + p.GetModel(), " | ", p.GetPrice())
		count++
	}
}

func (of *orderFacade) removeFromOrder(countTable int) {
	s := *of.products
	productListLength := len(s)
	for i, _ := range s {
		if countTable == i {
			s[productListLength-1], s[i] = s[i], s[productListLength-1]
		}
	}
	*of.products = s[:productListLength-1]

	//s := *of.products
	//s = append(s[:id], s[id+1:]...)
	//*of.products = s
}

func calculateTotalPrice(ps *[]models.Product) float32 {
	var totalPrice float32
	var totalDiscount float32

	for _, p := range *ps {
		if p.GetPrice() > 0 && p.GetPrice() < 10 {
			totalDiscount += 30.0
		} else if p.GetPrice() > 10 && p.GetPrice() < 100 {
			totalDiscount += 20.0
		} else if p.GetPrice() > 100 {
			totalDiscount += 10.0
		} else {
			totalDiscount += 5
		}

		if p.GetModel() == "S2" {
			totalDiscount += 30.0
		} else if p.GetModel() == "n320" {
			totalDiscount += 20.0
		} else if p.GetModel() == "Iphone" {
			totalDiscount += 10.0
		} else {
			totalDiscount += 5
		}

		if p.GetModel() == "Xiaomi" {
			totalDiscount += 30.0
		} else if p.GetModel() == "Samsung" {
			totalDiscount += 20.0
		} else if p.GetModel() == "Apple" {
			totalDiscount += 10.0
		} else{
			totalDiscount += 5
		}
		totalPrice = totalPrice + p.GetPrice()*(100 - totalDiscount)
	}

	return totalPrice
}

func (of *orderFacade) makeOrder(ps *[]models.Product){
	of.userRepo.RemoveMoneyFromBalance(of.user, calculateTotalPrice(ps))
	of.notification.sendSuccesOrderNotification()
}

func (of *orderFacade) closeOrder(){
	of.notification.sendCloseOrderNotification()
}

type notification struct {}

func (n *notification) sendSuccesOrderNotification() {
	fmt.Println("Order was made successfully!")
}

func (n *notification) sendCloseOrderNotification() {
	fmt.Println("Order was closed successfully!")
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