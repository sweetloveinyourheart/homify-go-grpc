package servers

import (
	"context"
	proto "homify-go-grpc/api/authentication"
	"homify-go-grpc/internal/authentication-service/services"
)

type GRPCAuthenticationServer struct {
	svc services.IAuthenticationService
	proto.UnimplementedAuthenticationServer
}

func NewGRPCAuthenticationServer() *GRPCAuthenticationServer {
	svc := services.NewAuthenticationService()

	return &GRPCAuthenticationServer{
		svc: svc,
	}
}

func (a *GRPCAuthenticationServer) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	success, err := a.svc.SignUp(req)

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
