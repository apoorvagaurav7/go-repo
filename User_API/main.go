package main

import (
	"context"
	"fmt"
	"go-repo/User_API/config"
	"go-repo/User_API/controllers"
	"go-repo/User_API/routes"
	"go-repo/User_API/services"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	ctx         context.Context
	server      *gin.Engine
)

func main() {
	server = gin.Default()
	InitializeDatabase()
	//InitializeAuthentication()
	InitializeProducts()
	InitializeUsers()
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer mongoClient.Disconnect(ctx1)
	server.Run(":6000")
}

func InitializeDatabase() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		log.Fatalf("Unable to connect to Database", err)
	} else {
		fmt.Println("Connected to Database")
	}
}
func InitializeProducts() {
	productCollection := config.GetCollection(mongoClient, "UserAPI", "TestProducts")
	productSvc := services.InitProductService(productCollection)
	productCtrl := controllers.InitProductController(productSvc)
	routes.ProductRoutes(server, *productCtrl)
}
func InitializeUsers() {
	userCollection := config.GetCollection(mongoClient, "UserAPI", "TestUsers")
	userService := services.InitUserService(userCollection)
	userCntrl := controllers.InitUserController(userService)
	routes.UserRoutes(server, *userCntrl)
}
