package services

import (
	"context"
	"errors"
	"fmt"
	"go-repo/User_API/entities"
	"go-repo/User_API/interfaces"
	"go-repo/User_API/utils"

	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	UserCollection *mongo.Collection
}

func InitUserService(collection *mongo.Collection) interfaces.IUser {
	return &UserService{UserCollection: collection}
}

func (uc *UserService) Register(user *entities.User) (*entities.SignupResponse, error) {
	ctx := context.Background()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	res, err := uc.UserCollection.InsertOne(ctx, &user)
	fmt.Println(res)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("email already exist")
		}
		return nil, err
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := uc.UserCollection.Indexes().CreateOne(ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser entities.SignupResponse
	query := bson.D{{"_id", res.InsertedID}}
	fmt.Println(res.InsertedID)

	err2 := uc.UserCollection.FindOne(ctx, query).Decode(&newUser)
	fmt.Println(uc.UserCollection.FindOne(ctx, query))

	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}

	return &newUser, nil
}

func (uc *UserService) Login(user *entities.Login) (*entities.LoginResponse, error) {
	ctx := context.Background()
	query := bson.M{"email": strings.ToLower(user.Email)}
	var loginResult *entities.User
	err := uc.UserCollection.FindOne(ctx, query).Decode(&loginResult)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//compare hashsed password with user entered password
	err2 := utils.VerifyPassword(loginResult.Password, user.Password)
	if err != nil {
		return nil, err2
	}
	// Generate the JWT token
	token, refreshToken, err := utils.GenerateAllTokens(loginResult.Email, loginResult.FirstName, loginResult.LastName, loginResult.ID.Hex())
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &entities.LoginResponse{Token: token, RefreshToken: refreshToken}, nil
}

// func (u *UserService) Logout(logoutRequest *entities.Logout) (*entities.LogoutResponse, error) {

//     response := &entities.LogoutResponse{
//         Message:"Logout successful",
//     }
//     return response, nil
// }
