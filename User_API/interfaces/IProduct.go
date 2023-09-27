package interfaces

import "github.com/apoorvagaurav7/GO-REPO/entities"

type IUser interface {
	Register(user *entities.User) (*entities.SignupResponse, error)
	Login(user *entities.Login) (*entities.LoginResponse, error)
	GetUser(userId string) (*entities.User, error)
	//GetUserDetails(email string) (*entities.SignupResponse,error)
	//Login(user *entities.User) string
	//  GetProfile(userId int) (*entities.User)
	//  SearchUser(searchQuery string)
	//  Logout(userId int) (string)
}
