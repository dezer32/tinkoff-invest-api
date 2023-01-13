package invest_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

var (
	baseUri = "/tinkoff.public.invest.invest_client.contract.v1.InstrumentsService"
)

type InstrumentsClient interface {
	Shares(ctx context.Context, req *SharesRequest, opts ...grpc.CallOption) (*SharesResponse, error)
}

type instrumentsClient struct {
	client grpc.ClientConnInterface
}

func NewInstrumentsClient(client grpc.ClientConnInterface) InstrumentsClient {
	return &instrumentsClient{client}
}

func (i *instrumentsClient) Shares(ctx context.Context, req *SharesRequest, opts ...grpc.CallOption) (*SharesResponse, error) {
	res := new(SharesResponse)
	err := i.client.Invoke(ctx, fmt.Sprintf("%s/%s", baseUri, "Shares"), req, res, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
