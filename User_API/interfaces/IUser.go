package interfaces

import "go-repo/User_API/entities"

type IUser interface {
	Register(user *entities.User) (*entities.SignupResponse, error)
	Login(user *entities.Login) (*entities.LoginResponse, error)
	//Logout(*entities.Logout) (*entities.LogoutResponse, error)
}
