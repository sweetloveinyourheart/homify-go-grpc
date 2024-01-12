package authentication_service

type Client struct {
	Endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

// func NewGRPCClient(remoteAddr string) (proto.PriceFetcherClient, error) {
// 	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}

// 	c := proto.NewPriceFetcherClient(conn)

// 	return c, nil
// }
