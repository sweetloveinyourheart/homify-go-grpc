package grpc_client

import (
	proto "homify-go-grpc/api/authentication"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCAuthenticationClient(remoteAddr string) (proto.AuthenticationClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := proto.NewAuthenticationClient(conn)

	return c, nil
}
