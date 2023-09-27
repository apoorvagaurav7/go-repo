package services

import (
	"context"
	"fmt"
	"go-repo/User_API/entities"
	"go-repo/User_API/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct {

	return &ProductService{Product: collection}
}

func (p *ProductService) AddProducts(product *entities.Product) (string, error) {
	product.ID = primitive.NewObjectID()
	_, err := p.Product.InsertOne(context.Background(), product)
	if err != nil {
		return "", err
	} else {
		return "Record Inserted Successfully", nil
	}
}
func (p *ProductService) GetProductsById() ([]*entities.Product, error) {
	result, err := p.Product.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		//do something
		fmt.Println(result)
		//build the array of products for the cursor that we received.
		var products []*entities.Product
		for result.Next(context.TODO()) {
			product := &entities.Product{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}
			//fmt.Println(product)
			products = append(products, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			return []*entities.Product{}, nil
		}
		return products, nil
	}

}
func (p *ProductService) SearchProduct(product *entities.Product) (string, error) {

	searchQuery := bson.M{"name": bson.M{"$regex": "keyword-to-search", "$options": "i"}}
	_, err := p.Product.Find(context.Background(), searchQuery)
	if err != nil {
		return "", err
	} else {
		return "Product found Successfully", nil
	}
}
