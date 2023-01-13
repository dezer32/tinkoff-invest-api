package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dezer32/tinkoff-invest-api/configs"
	"github.com/dezer32/tinkoff-invest-api/internal/config"
	"github.com/dezer32/tinkoff-invest-api/internal/generated/investapi"
	"github.com/dezer32/tinkoff-invest-api/pkg/client"
	"log"
	"os"
	"time"
)

var (
	cfg *configs.Config
	c   *client.Client
)

func init() {
	var err error
	configFiles := []string{
		"configs/client.yaml",
		"configs/time.yaml",
	}
	cfg, err = config.Load(configFiles...)
	if err != nil {
		log.Fatalf("%s : when load configs", err)
	}

	c, err = client.New(cfg)
	if err != nil {
		log.Fatalf("%s : when connect to api", err)
	}
}

func main() {

	//
	req := &investapi.InstrumentsRequest{
		InstrumentStatus: 0,
	}
	shares, err := c.Services.Instruments.Shares(context.Background(), req)
	if err != nil {
		log.Fatalf("%s : when load shares", err)
	}

	res := make(map[string]string)

	for _, share := range shares.Instruments {
		if _, ok := res[share.Figi]; ok != true {
			res[share.Figi] = share.Name
			continue
		}

		log.Printf("Перезаписал FIGI (%s) с Shares..", share.Figi)
	}

	data, _ := json.Marshal(res)
	fileName := fmt.Sprintf("figi.%d.txt", time.Now().Unix())
	os.WriteFile(fileName, data, os.ModePerm)
}