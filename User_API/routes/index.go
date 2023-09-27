package routes

import (
	"go-repo/User_API/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, a controllers.UserController) {
	user := r.Group("/userapi/user")
	user.POST("/register", a.Register)
	user.POST("/login", a.Login)
	//user.POST("/logout")
}

func ProductRoutes(r *gin.Engine, p controllers.ProductController) {
	product := r.Group("/userapi/product")
	product.POST("/addproduct", p.AddProduct)
	product.GET("/getproductsbyid", p.GetProductsById)
	product.GET("/searchproduct", p.SearchProduct)
}
