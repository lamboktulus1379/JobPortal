package usecase

import (
	"context"
	"crypto/md5"
	"fmt"
	"jobportal/domain/dto"
	"jobportal/domain/model"
	"jobportal/domain/repository"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type IUserUsecase interface {
	Login(ctx context.Context, req model.ReqLogin) dto.ResLogin
}

type UserUsecase struct {
	userRepository repository.IUser
}

func NewUserUsecase(userRepository repository.IUser) IUserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (userUsecase *UserUsecase) Login(ctx context.Context, req model.ReqLogin) dto.ResLogin {
	var res dto.ResLogin

	user, err := userUsecase.userRepository.GetByUserName(ctx, req.UserName)
	log.Printf("Username: %s\n", user.UserName)
	if err != nil {
		log.Printf("User not found. %v\n", err)
		res.ResponseCode = "401"
		res.ResponseMessage = "Unautorized."
		return res
	}
	md5Req := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))

	if md5Req != user.Password {
		res.ResponseCode = "401"
		res.ResponseMessage = "Unautorized."
		return res
	}

	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println("Secret Key: ", secretKey)
	mySigningKey := []byte(secretKey)

	// Create the Claims
	expiration := time.Now().Add(5 * time.Minute)

	claims := model.UserClaims{
		UserName: req.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			Issuer:    fmt.Sprint(user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", accessToken, err)
	if err != nil {
		res.ResponseCode = "401"
		res.ResponseMessage = "Unautorized"
		return res
	}

	res.ResponseCode = "200"
	res.ResponseMessage = "Success"
	res.Data.AccessToken = accessToken
	res.Data.ExpiresAt = expiration.Unix()

	return res
}
