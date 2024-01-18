package server

import (
	"context"
	proto "homify-go-grpc/api/authentication"
	"homify-go-grpc/internal/authentication-service/services"

	"gorm.io/gorm"
)

type GRPCAuthenticationServer struct {
	svc services.IAuthenticationService
	proto.UnimplementedAuthenticationServer
}

func NewGRPCAuthenticationServer(db *gorm.DB) *GRPCAuthenticationServer {
	svc := services.NewAuthenticationService(db)

	return &GRPCAuthenticationServer{
		svc: svc,
	}
}

func (a *GRPCAuthenticationServer) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	registerData := services.RegisterAccount{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		Gender:   req.Gender,
		Birthday: req.Birthday,
		Phone:    req.Phone,
	}

	success, err := a.svc.SignUp(registerData)

	if err != nil {
		return &proto.SignUpResponse{
			Message: err.Error(),
			Success: success,
		}, nil
	}

	return &proto.SignUpResponse{
		Message: "New account was created",
		Success: success,
	}, nil
}

func (a *GRPCAuthenticationServer) SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error) {
	accountData := services.LoginAccount{
		Email:    req.Email,
		Password: req.Password,
	}

	tokens, err := a.svc.SignIn(accountData)

	if err != nil {
		return nil, err
	}

	return &proto.SignInResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (a *GRPCAuthenticationServer) VerifyJwtToken(ctx context.Context, req *proto.VerifyJwtTokenRequest) (*proto.VerifyJwtTokenResponse, error) {
	token := req.Token

	authenticatedUser, err := a.svc.VerifyJwtToken(token)
	if err != nil {
		return &proto.VerifyJwtTokenResponse{}, err
	}

	return &proto.VerifyJwtTokenResponse{
		UserId: int32(authenticatedUser.Id),
		Email:  authenticatedUser.Email,
	}, nil
}
