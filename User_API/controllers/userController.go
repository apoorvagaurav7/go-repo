package controllers

import (
	"fmt"
	"go-repo/User_API/entities"
	"go-repo/User_API/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService interfaces.IUser
}

func InitUserController(userService interfaces.IUser) *UserController {
	return &UserController{UserService: userService}
}

func (a *UserController) Register(c *gin.Context) {
	fmt.Println("User Registered")
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.UserService.Register(&user)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (a *UserController) Login(c *gin.Context) {
	fmt.Println("User LoggedIn")
	var user entities.Login
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.UserService.Login(&user)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

// func (ulc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "Logout successful")
// }
