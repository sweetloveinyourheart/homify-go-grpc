package grpc_client

import (
	proto "homify-go-grpc/api/property"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCPropertyClient(remoteAddr string) (proto.PropertyClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := proto.NewPropertyClient(conn)

	return c, nil
}
