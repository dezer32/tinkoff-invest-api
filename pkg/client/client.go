package client

import (
	"crypto/tls"
	"fmt"
	"github.com/dezer32/tinkoff-invest-api/configs"
	"github.com/dezer32/tinkoff-invest-api/pkg/generated/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	Connection *grpc.ClientConn
	Services   *services
}

type services struct {
	Instruments investapi.InstrumentsServiceClient
	MarketData  investapi.MarketDataServiceClient
}

func New(cfg *configs.Config) (client *Client, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName: cfg.Client.Endpoint.URL,
		})),
		grpc.WithPerRPCCredentials(TokenAuth{
			Token: cfg.Client.Token,
		}),
	}

	endpoint := fmt.Sprintf("%s:%d", cfg.Client.Endpoint.URL, cfg.Client.Endpoint.Port)
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return
	}

	client = &Client{
		Connection: conn,
		Services: &services{
			Instruments: investapi.NewInstrumentsServiceClient(conn),
			MarketData:  investapi.NewMarketDataServiceClient(conn),
		},
	}

	return
}
