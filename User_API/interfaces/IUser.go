package interfaces

import "github.com/apoorvagaurav7/GO-REPO/entities"

type IProduct interface {
	Insert(product *entities.Product) (string, error)
	GetProducts() ([]*entities.Product, error)
	UpdateProduct(product *entities.Product) (string, error)
}
