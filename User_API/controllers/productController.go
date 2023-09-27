package controllers

import (
	"fmt"
	"go-repo/User_API/entities"
	"go-repo/User_API/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService interfaces.IProduct
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) AddProduct(c *gin.Context) {
	fmt.Println("Invoked controller")
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		return
	}
	result, err := p.ProductService.AddProducts(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}

	//extract the product from the request Object
	//call the service.insert
	//p.ProductService.InsertOne()
}

func (p ProductController) GetProductsById(c *gin.Context) {
	result, err := p.ProductService.GetProductsById()
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (p ProductController) SearchProduct(c *gin.Context) {
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		return
	}
	fmt.Println(product.ID)
	result, err := p.ProductService.SearchProduct(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}
