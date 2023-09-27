package interfaces

import "go-repo/User_API/entities"

type IProduct interface {
	AddProducts(product *entities.Product) (string, error)
	GetProductsById() ([]*entities.Product, error)
	SearchProduct(product *entities.Product) (string, error)
}
