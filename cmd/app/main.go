package main

import (
	"context"
	"github.com/dezer32/tinkoff-invest-api/internal/config"
	invest_client2 "github.com/dezer32/tinkoff-invest-api/internal/invest_client"
	"github.com/dezer32/tinkoff-invest-api/pkg/invest_client"
	"log"
)

func main() {
	cfg, err := config.Load("configs/client.yaml")
	if err != nil {
		log.Fatalf("%s : when load configs", err)
	}

	c, err := invest_client.New(cfg)
	if err != nil {
		log.Fatalf("%s : when connect to api")
	}

	req := &invest_client2.SharesRequest{InstrumentStatus: invest_client2.InstrumentStatusAll}
	resp, err := c.InstrumentsClient.Shares(context.Background(), req)
	if err != nil {
		log.Fatalf("%s : when load instruments", err)
	}

	log.Printf("Result:\n%v", resp)
}
