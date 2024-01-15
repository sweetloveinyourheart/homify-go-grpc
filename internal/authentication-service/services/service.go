package services

import (
	"database/sql"
	"fmt"
	proto "homify-go-grpc/api/authentication"
	"homify-go-grpc/internal/authentication-service/database"
	"homify-go-grpc/internal/authentication-service/models"
	"homify-go-grpc/internal/authentication-service/repositories"
	"homify-go-grpc/internal/authentication-service/utils"
)

type IAuthenticationService interface {
	SignUp(req *proto.SignUpRequest) (bool, error)
}

type AuthenticationService struct {
	accountRepository repositories.IAccountRepository
	userRepository    repositories.IUserRepository
}

func NewAuthenticationService() IAuthenticationService {
	db := database.GetDB()

	return &AuthenticationService{
		accountRepository: repositories.NewAccountRepository(db),
		userRepository:    repositories.NewUserRepository(db),
	}
}

func (a *AuthenticationService) SignUp(newUserData *proto.SignUpRequest) (bool, error) {
	existingAccount, getAccountErr := a.accountRepository.GetAccountByField("Email", newUserData.Email)
	if existingAccount != nil || getAccountErr != nil {
		return false, fmt.Errorf("account is already exist")
	}

	// Create new account
	hashedPassword, hashErr := utils.HashPassword(newUserData.Password)
	if hashErr != nil {
		return false, hashErr
	}

	newAccount := models.Account{
		Email:    newUserData.Email,
		Password: hashedPassword,
	}

	createAccountErr := a.accountRepository.CreateAccount(&newAccount)
	if createAccountErr != nil {
		return false, createAccountErr
	}

	newUser := models.User{
		Account:  newAccount,
		FullName: newUserData.FullName,
		Birthday: sql.NullString{String: newUserData.Birthday, Valid: true},
		Gender:   sql.NullString{String: newUserData.Gender, Valid: true},
		Phone:    sql.NullString{String: newUserData.Phone, Valid: true},
	}

	createUserErr := a.userRepository.CreateUser(&newUser)
	if createUserErr != nil {
		return false, createUserErr
	}

	return true, nil
}
