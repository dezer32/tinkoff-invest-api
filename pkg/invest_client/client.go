package invest_client

import (
	"crypto/tls"
	"fmt"
	"github.com/dezer32/tinkoff-invest-api/configs"
	"github.com/dezer32/tinkoff-invest-api/internal/invest_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	Connection        *grpc.ClientConn
	InstrumentsClient invest_client.InstrumentsClient
}

func New(cfg *configs.Config) (client *Client, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName: cfg.Client.Endpoint.URL,
		})),
		grpc.WithPerRPCCredentials(invest_client.TokenAuth{
			Token: cfg.Client.Token,
		}),
	}

	endpoint := fmt.Sprintf("%s:%d", cfg.Client.Endpoint.URL, cfg.Client.Endpoint.Port)
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return
	}

	client = &Client{
		Connection:        conn,
		InstrumentsClient: invest_client.NewInstrumentsClient(conn),
	}

	return
}
