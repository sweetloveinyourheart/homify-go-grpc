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
	return &proto.SignUpResponse{
		Message: "Success",
	}, nil
}
