package chainOfResponsobility

import "GolangFinalProject/models"

type storeDepartment interface {
	execute(*models.Product)
	setNext(storeDepartment)
}

type  struct {
	
}
