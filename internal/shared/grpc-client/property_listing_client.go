package grpc_client

import (
	proto "homify-go-grpc/api/property-listing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCPropertyListingClient(remoteAddr string) (proto.PropertyListingClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := proto.NewPropertyListingClient(conn)

	return c, nil
}
