package services

import (
	"database/sql"
	"fmt"
	"homify-go-grpc/internal/authentication-service/configs"
	"homify-go-grpc/internal/authentication-service/models"
	"homify-go-grpc/internal/authentication-service/repositories"
	"homify-go-grpc/internal/authentication-service/utils"
	"log"
	"time"

	"gorm.io/gorm"
)

type LoginAccount struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type RegisterAccount struct {
	Email    string
	Password string
	FullName string
	Gender   string
	Birthday string
	Phone    string
}

type IAuthenticationService interface {
	SignUp(userData RegisterAccount) (bool, error)
	SignIn(userData LoginAccount) (Tokens, error)
}

type AuthenticationService struct {
	accountRepository repositories.IAccountRepository
	userRepository    repositories.IUserRepository
}

func NewAuthenticationService(db *gorm.DB) IAuthenticationService {
	return &AuthenticationService{
		accountRepository: repositories.NewAccountRepository(db),
		userRepository:    repositories.NewUserRepository(db),
	}
}

func (a *AuthenticationService) SignUp(newUserData RegisterAccount) (bool, error) {
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

func (a *AuthenticationService) SignIn(userData LoginAccount) (Tokens, error) {
	account, getAccountErr := a.accountRepository.GetAccountByField("Email", userData.Email)
	if getAccountErr != nil {
		return Tokens{}, getAccountErr
	}

	if account == nil {
		return Tokens{}, fmt.Errorf("account is not found")
	}

	isValidPassword := utils.CheckPasswordHash(userData.Password, account.Password)
	if !isValidPassword {
		return Tokens{}, fmt.Errorf("password is not valid")
	}

	payload := utils.TokenPayload{
		Id:    account.ID,
		Email: account.Email,
	}

	configurations := configs.GetConfig()

	// Generate Token
	accessToken, accessTokenErr := utils.GenerateToken(payload, configurations.JwtSecret, 15*time.Minute)
	if accessTokenErr != nil {
		log.Println(accessTokenErr.Error())
		return Tokens{}, fmt.Errorf("email or password is not valid")
	}

	refreshToken, refreshTokenErr := utils.GenerateToken(payload, configurations.JwtSecret, 24*time.Hour)
	if refreshTokenErr != nil {
		log.Println(refreshTokenErr.Error())
		return Tokens{}, fmt.Errorf("email or password is not valid")
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
