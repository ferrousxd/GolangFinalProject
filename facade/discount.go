package facade

import (
	"GolangFinalProject/models"
	"GolangFinalProject/repositories"
	"fmt"
)

type discountFacade struct {
	price float64
	model string
	company string
}

func newDiscountFacade(product models.Product) *discountFacade  {
	productPrice := product.GetPrice()
	productModel := product.GetModel()
	productCompany := product.GetCompany()
	fmt.Println("Starting to run discount...")
	discountFacde := &discountFacade{
		price: discountByPrice(productPrice),
		model: discountByModel(productModel),
		company: discountByCompany(productCompany),
	}
	fmt.Println("Discount is on ran")
}

type NewPrice struct {
	
}