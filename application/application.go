package application

import (
	"GolangFinalProject/models"
	"GolangFinalProject/repositories"
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Application struct {
	productRepo repositories.ProductRepository
	userRepo repositories.UserRepository
}

var user *models.User

func NewApplication(productRepo repositories.ProductRepository,
	userRepo repositories.UserRepository) *Application  {
	return &Application{
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (a *Application) Start()  {
	for {
		fmt.Println("Welcome to our phone shop. What would you like to do?")
		fmt.Println("1. Sign in")
		fmt.Println("2. Sign up")
		fmt.Println("3. Close application (press any key to exit)")

		var choice int

		fmt.Scan(&choice)

		if choice == 1 {
			a.SignIn()
		} else if choice == 2 {
			a.SignUp()
		}else {
			break
		}
	}
}

func (a *Application) SignIn()  {
	for {
		fmt.Println("Enter your username and password")

		var username string
		var password string

		fmt.Scan(&username)
		fmt.Scan(&password)

		user = a.userRepo.GetUserByLogin(username, password)

		if user != nil {
			fmt.Println("Successfully signed in!")

			if user.GetRole() == "User" {
				a.MainMenu()
				break
			} else if user.GetRole() == "Admin" {
				a.AdminMenu()
				break
			}
		} else {
			fmt.Println("Wrong credentials! What you want to do?")
			fmt.Println("1. Try again.")
			fmt.Println("2. Exit.")

			var choice int

			fmt.Scan(&choice)

			if choice != 1 {
				break
			}
		}
	}
}

func (a *Application) SignUp()  {
	for  {
		fmt.Println("Please, enter username, email and password:")

		var username string
		var email string
		var password string

		fmt.Scan(&username)
		fmt.Scan(&email)
		fmt.Scan(&password)

		if !isValidEmail(email) {
			fmt.Println("Not valid email format! What would you like to do?")
			fmt.Println("1. Try again.")
			fmt.Println("2. Exit.")

			var choice int

			fmt.Scan(&choice)

			if choice != 1 {
				break
			}
		} else {
			newUserBuilder := models.UserBuilder{}

			newUser := *newUserBuilder.
				SetUsername(username).
				SetEmail(email).
				SetPassword(password).
				Build()

			a.userRepo.InsertUser(newUser)

			fmt.Println("User successfully signed up!")
			break
		}
	}
}

func (a *Application) MainMenu()  {
	for {
		fmt.Println("Main menu.")
		fmt.Println("1. Make Order")
		fmt.Println("2. Get list of products")
		fmt.Println("3. Get list of subscribed products")
		fmt.Println("4. Subscribe to product updates")
		fmt.Println("5. Unsubscribe from product updates")
		fmt.Println("6. Log out")

		var choice int

		fmt.Scan(&choice)

		if choice == 1 {
			// Азаткали, пиши свой код здесь


			


		} else if choice == 2 {
			fmt.Println("List of products:")

			products := a.productRepo.GetAllProducts()

			a.printSliceOfProducts(products)
		} else if choice == 3 {
			fmt.Println("List of subscribed products:")

			products := a.productRepo.GetProductsBySubscriberId(user.GetId())

			a.printSliceOfProducts(products)
		} else if choice == 4 {
			fmt.Println("Enter the ID of the product, that you want to subscribe for:")

			products := a.productRepo.GetAllProducts()

			a.printSliceOfProducts(products)

			var productId int

			fmt.Scan(&productId)

			a.userRepo.ChangeSubscriptionStatus(productId, user.GetId(), "add")
		} else if choice == 5 {
			fmt.Println("Enter the ID of the product, that you want to unsubscribe from:")

			products := a.productRepo.GetProductsBySubscriberId(user.GetId())

			a.printSliceOfProducts(products)

			var productId int

			fmt.Scan(&productId)

			a.userRepo.ChangeSubscriptionStatus(productId, user.GetId(), "remove")
		} else {
			user = nil
			break
		}
	}
}

func (a *Application) AdminMenu()  {
	for {
		fmt.Println("Admin Menu:")
		fmt.Println("1. Add new product")
		fmt.Println("2. Delete product")
		fmt.Println("3. Get list of products")
		fmt.Println("4. Notify users about products")
		fmt.Println("5. Log out")

		var choice int

		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Please, enter model, company and price for new product:")

			var model string
			var company string
			var price float32

			fmt.Scan(&model)
			fmt.Scan(&company)
			fmt.Scan(&price)

			newProductBuilder := models.ProductBuilder{}

			newProduct := *newProductBuilder.
				SetModel(model).
				SetCompany(company).
				SetPrice(price).
				Build()

			a.productRepo.InsertProduct(newProduct)

			fmt.Println("New product was added!")
		} else if choice == 2 {
			fmt.Println("Please, enter id of product which you want to delete:")

			var productId int

			fmt.Scan(&productId)

			a.productRepo.DeleteProduct(productId)

			fmt.Println("Product was deleted!")
		} else if choice == 3 {
			fmt.Println("List of products:")

			products := a.productRepo.GetAllProducts()

			a.printSliceOfProducts(products)
		} else if choice == 4 {
			fmt.Println("Choose the product that you want to notify users about")

			products := a.productRepo.GetAllProducts()

			a.printSliceOfProducts(products)

			var productId int

			fmt.Scan(&productId)

			product := a.productRepo.GetProductById(productId)

			subscribers := a.userRepo.GetSubscribersByProductId(productId)

			for _, o := range subscribers {
				product.AddObserver(o)
			}

			product.NotifyAllObservers()
		} else {
			user = nil
			break
		}
	}
}


func (a *Application) printSliceOfProducts(products []*models.Product) {
	for _, p := range products {
		fmt.Println(p.GetId(), " | ", p.GetCompany(), " | " + p.GetModel(), " | ", p.GetPrice())
	}
}

func isValidEmail(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}

	return emailRegex.MatchString(e)
}