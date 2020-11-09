package application

import (
	"GolangFinalProject/facade"
	"GolangFinalProject/models"
	"GolangFinalProject/repositories"
	"fmt"
	"regexp"
)

// Regex pattern for email
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Application struct, that stores all possible repository objects
type Application struct {
	productRepo repositories.ProductRepository
	userRepo repositories.UserRepository
}

// user variable, that stores everything about the logged in user, except his password
var user *models.User
// selectedProducts variable, that stores the slice of selected products
var selectedProducts []models.Decorator

// Constructor for Application
func NewApplication(productRepo repositories.ProductRepository,
	userRepo repositories.UserRepository) *Application  {
	return &Application{
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

// Starting the application
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
		} else {
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

		// Initializing user object with GetUserByLogin() method
 		user = a.userRepo.GetUserByLogin(username, password)

		if user != nil {
			fmt.Println("Successfully signed in!")

			if user.GetRole() == "User" {
				// If user role is equal to "User" - it redirects you to standard Main Menu
				a.MainMenu()
				break
			} else if user.GetRole() == "Admin" {
				// If user role is equal to "Admin" - it redirects you to special Admin Menu
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
	for {
		fmt.Println("Please, enter username, email and password:")

		var username string
		var email string
		var password string

		fmt.Scan(&username)
		fmt.Scan(&email)
		fmt.Scan(&password)

		// Here we check email format is valid or not
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
			// Creating newUserBuilder object, for building the User object
			newUserBuilder := models.UserBuilder{}

			// Creating newUser with builder
			newUser := *newUserBuilder.
				SetUsername(username).
				SetEmail(email).
				SetPassword(password).
				Build()

			// Inserting our new user into database
			a.userRepo.InsertUser(newUser)

			fmt.Println("User successfully signed up!")
			break
		}
	}
}

func (a *Application) MainMenu()  {
	for {
		fmt.Println("Main menu.")
		fmt.Println("1. Order menu")
		fmt.Println("2. Get list of products")
		fmt.Println("3. Get list of subscribed products")
		fmt.Println("4. Subscribe to product updates")
		fmt.Println("5. Unsubscribe from product updates")
		fmt.Println("6. Replenish balance")
		fmt.Println("7. Check balance")
		fmt.Println("8. Log out")

		var choice int

		fmt.Scan(&choice)

		if choice == 1 {
			// Азаткали, пиши свой код здесь
			a.OrderMenu()
		} else if choice == 2 {
			fmt.Println("List of products:")

			// Creating a slice of all available products
			products := a.productRepo.GetAllProducts()

			// Printing out the products with special method
			a.printSliceOfProducts(products)
		} else if choice == 3 {
			fmt.Println("List of subscribed products:")

			// Creating a slice of products, that user is observing
			products := a.productRepo.GetProductsBySubscriberId(user.GetId())

			// Printing out the products with special method
			a.printSliceOfProducts(products)
		} else if choice == 4 {
			fmt.Println("Enter the ID of the product, that you want to subscribe for:")

			// Creating a slice of all available products
			products := a.productRepo.GetAllProducts()

			// Printing out the products with special method
			a.printSliceOfProducts(products)

			var productId int

			fmt.Scan(&productId)

			// Adding the product to subscriptions for the signed in user
			a.userRepo.ChangeSubscriptionStatus(productId, user.GetId(), "add")
		} else if choice == 5 {
			fmt.Println("Enter the ID of the product, that you want to unsubscribe from:")

			// Creating a slice of all available products
			products := a.productRepo.GetProductsBySubscriberId(user.GetId())

			// Printing out the products with special method
			a.printSliceOfProducts(products)

			var productId int

			fmt.Scan(&productId)

			// Deleting the specific product for subscriptions for the signed in user
			a.userRepo.ChangeSubscriptionStatus(productId, user.GetId(), "remove")
		} else if choice == 6 {
			fmt.Println("Enter the sum which you want to add:")

			var amount float32

			fmt.Scan(&amount)

			// Updating (or adding) money to user balance
			a.userRepo.AddMoneyToBalance(user, amount)
			// Getting the userId from user object
			userId := user.GetId()
			// Rewriting the original user object by receiving the new object with GetUserById()
			user = a.userRepo.GetUserById(userId)

			fmt.Println("Balance was updated successfully!")
		} else if choice == 7 {
			// Printing out the balance of the user
			fmt.Println("Your current balance: ", user.GetBalance())
		} else {
			// Destroying the user object
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

			// Creating newProductBuilder for creating product
			newProductBuilder := models.ProductBuilder{}

			// Creating newProduct object
			newProduct := *newProductBuilder.
				SetModel(model).
				SetCompany(company).
				SetPrice(price).
				Build()

			// Adding the newProduct to our database
			a.productRepo.InsertProduct(newProduct)

			fmt.Println("New product was added!")
		} else if choice == 2 {
			fmt.Println("Please, enter ID of product which you want to delete:")

			var productId int

			fmt.Scan(&productId)

			// Deleting the product from our database by productId
			a.productRepo.DeleteProduct(productId)

			fmt.Println("Product was deleted!")
		} else if choice == 3 {
			fmt.Println("List of products:")

			// Creating a slice of all available products
			products := a.productRepo.GetAllProducts()

			// Printing out the products with special method
			a.printSliceOfProducts(products)
		} else if choice == 4 {
			fmt.Println("Choose the product that you want to notify users about")

			// Creating a slice of all available products for notifying the users
			products := a.productRepo.GetAllProducts()

			// Printing out the products with special method
			a.printSliceOfProducts(products)

			var productId int

			fmt.Scan(&productId)

			// Receiving the product object with GetProductById() method
 			product := a.productRepo.GetProductById(productId)

 			// Getting slice of observers by productId
			subscribers := a.userRepo.GetSubscribersByProductId(productId)

			// Filling the observerList field in observable
			for _, o := range subscribers {
				product.AddObserver(o)
			}

			// Notifying all observers
			product.NotifyAllObservers()
		} else {
			// Destroying the user
			user = nil
			break
		}
	}
}

func (a *Application) OrderMenu() {
	orderFacade := facade.NewOrderFacade(user, selectedProducts, a.userRepo)
	for {
		fmt.Println("Order menu:")
		fmt.Println("1. Add product to cart")
		fmt.Println("2. Remove product from cart")
		fmt.Println("3. Get list of ordered products")
		fmt.Println("4. Make order")
		fmt.Println("5. Cancel order")
		fmt.Println("6. Exit")

		var choice int

		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Enter the ID of the product, that you want to add to cart:")

			// Creating a slice of all available products for making the order
			products := a.productRepo.GetAllProducts()

			// Printing out the products with special method
			a.printSliceOfProducts(products)

			var productId int

			fmt.Scan(&productId)

			// Receiving the product object with GetProductById() method
			product := a.productRepo.GetProductById(productId)

			// The base model has 64 GB of memory. User can modify his Phone with decorator pattern
			fmt.Println("This", product.GetModel(), "has 64 GB of memory. Do you want to change the amount of storage? (Yes / No)")

			// Here we store the option for memory capacity (if user is satisfied with base model characteristics)
			var choiceMemory string

			fmt.Scan(&choiceMemory)

			// Variable that stores the product decorator for future modifications
			var productDifMemory models.Decorator

			if choiceMemory == "Yes" {
				// Here user should choose the desired amount of memory
				fmt.Println("Choose amount of memory:")
				fmt.Println("1. 128 GB")
				fmt.Println("2. 256 GB")

				var choiceOfMemory int

				fmt.Scan(&choiceOfMemory)

				if choiceOfMemory == 1 {
					// Decorating our base product with 128GB modification
					productDifMemory = &models.With128GB{Decorator: product}
				} else if choiceOfMemory == 2 {
					// Decorating our base product with 256GB modification
					productDifMemory = &models.With256GB{Decorator: product}
				}
			} else if choiceMemory == "No" {
				// The base model is selected automatically, if user types "No"
				productDifMemory = product
			}

			fmt.Println("Do you want product with case? (Yes / No)")

			// The user has choice: if he wants to have the phone with case
			var choiceCase string

			fmt.Scan(&choiceCase)

			if choiceCase == "Yes" {
				// Case decorator for our productDifMemory
				productWithCase := models.WithCase{Decorator: productDifMemory}

				// Appending our modified phone to selectedProducts
				selectedProducts = append(selectedProducts, &productWithCase)
				// Overriding facades Products field
				orderFacade.Products = selectedProducts

				fmt.Println("Product was successfully added to cart!")
			} else if choiceCase == "No" {
				// Appending our modified phone to selectedProducts
				selectedProducts = append(selectedProducts, productDifMemory)
				// Overriding facades Products field
				orderFacade.Products = selectedProducts

				fmt.Println("Product was successfully added to cart!")
			}
		} else if choice == 2 {
			fmt.Println("Enter the row number of the product, that you want to remove from cart:")

			// Printing out all products from order, if user desires to delete one of the products from his cart
			orderFacade.PrintProduct()

			var rowNumber int

			fmt.Scan(&rowNumber)

			// Removing the product from order by rowNumber
			orderFacade.RemoveFromOrder(rowNumber)

			fmt.Println("Product was successfully removed from cart!")
		} else if choice == 3 {
			fmt.Println("List of added products:")

			/// Printing out all products from order
			orderFacade.PrintProduct()
		} else if choice == 4 {
			// Confirming the order
			orderFacade.MakeOrder()
			// Overriding the user object
			user = a.userRepo.GetUserById(user.GetId())
		} else if choice == 5 {
			// Canceling the order
			orderFacade.CancelOrder()
			// Destroying the selected products
			selectedProducts = nil
		} else {
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