package client

import (
	"github.com/dezer32/tinkoff-invest-api/internal/generated/investapi"
	"github.com/dezer32/tinkoff-invest-api/internal/structs"
	"google.golang.org/grpc"
)

type Services struct {
	Instruments *Instruments
}

func NewServices(conn grpc.ClientConnInterface) *Services {
	return &Services{
		Instruments: NewInstruments(investapi.NewInstrumentsServiceClient(conn), structs.NewStructMapper()),
	}
}
