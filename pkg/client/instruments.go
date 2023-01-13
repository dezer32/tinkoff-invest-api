package client

import (
	"context"
	"github.com/dezer32/tinkoff-invest-api/internal/generated/investapi"
	"github.com/dezer32/tinkoff-invest-api/internal/structs"
)

type Instruments struct {
	Client investapi.InstrumentsServiceClient
	Mapper *structs.StructMapper
}

func NewInstruments(client investapi.InstrumentsServiceClient, mapper *structs.StructMapper) *Instruments {
	return &Instruments{
		Client: client,
		Mapper: mapper,
	}
}

func (i *Instruments) Shares(request *structs.SharesRequest) (response *structs.SharesResponse, err error) {
	resp, err := i.Client.Shares(context.Background(), &investapi.InstrumentsRequest{
		InstrumentStatus: investapi.InstrumentStatus(request.InstrumentStatus),
	})
	if err != nil {
		return
	}

	err = i.Mapper.Map(resp, &response)
	if err != nil {
		return
	}

	return
}
