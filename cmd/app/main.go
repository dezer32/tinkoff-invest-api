package main

import (
	"github.com/dezer32/tinkoff-invest-api/internal/config"
	"github.com/dezer32/tinkoff-invest-api/internal/structs"
	"github.com/dezer32/tinkoff-invest-api/pkg/client"
	"log"
)

func main() {
	cfg, err := config.Load("configs/client.yaml")
	if err != nil {
		log.Fatalf("%s : when load configs", err)
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatalf("%s : when connect to api")
	}
	//
	req := &structs.SharesRequest{InstrumentStatus: 0}
	resp, err := c.Services.Instruments.Shares(req)
	if err != nil {
		log.Fatalf("%s : when load instruments", err)
	}

	log.Printf("Result:\n%v", resp)
}
