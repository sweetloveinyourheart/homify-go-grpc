package services

import (
	proto "homify-go-grpc/api/authentication"
)

type IAuthenticationService interface {
	SignUp(req *proto.SignUpRequest) (bool, error)
}

type AuthenticationService struct{}

func NewAuthenticationService() IAuthenticationService {
	return &AuthenticationService{}
}

func (a *AuthenticationService) SignUp(req *proto.SignUpRequest) (bool, error) {
	return true, nil
}
