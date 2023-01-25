package client

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dezer32/tinkoff-invest-api/pkg/generated/investapi"
)

type Client struct {
	token string
	Url   string

	Connection *grpc.ClientConn
	Services   *services
}

type services struct {
	Instruments investapi.InstrumentsServiceClient
	MarketData  investapi.MarketDataServiceClient
}

func New(token, url string) (client *Client, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(
				&tls.Config{
					ServerName: url,
				},
			),
		),
		grpc.WithPerRPCCredentials(
			TokenAuth{
				Token: token,
			},
		),
	}

	endpoint := fmt.Sprintf("%s:443", url)
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return
	}

	client = &Client{
		token: token,
		Url:   url,

		Connection: conn,
		Services: &services{
			Instruments: investapi.NewInstrumentsServiceClient(conn),
			MarketData:  investapi.NewMarketDataServiceClient(conn),
		},
	}

	return
}
