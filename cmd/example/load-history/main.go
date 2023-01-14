package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dezer32/tinkoff-invest-api/configs"
	"github.com/dezer32/tinkoff-invest-api/pkg/client"
	"github.com/dezer32/tinkoff-invest-api/pkg/config"
	investapi2 "github.com/dezer32/tinkoff-invest-api/pkg/generated/investapi"
	"github.com/dezer32/tinkoff-invest-api/pkg/helpers"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"time"
)

var (
	cfg      *configs.Config
	c        *client.Client
	location *time.Location
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

	location, err = time.LoadLocation(cfg.Time.Location)
	if err != nil {
		log.Fatalf("%s : when load location", err)
	}
}

func main() {
	req := new(investapi2.GetCandlesRequest)
	req.Figi = "BBG000BBQCY0"
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, location)
	req.From = timestamppb.New(startDay)
	endDay := time.Date(now.Year(), now.Month(), now.Day()-1, 23, 59, 59, 0, location)
	req.To = timestamppb.New(endDay)
	req.Interval = 1

	candles, err := c.Services.MarketData.GetCandles(context.Background(), req)
	if err != nil {
		log.Fatalf("%s : when load candles", err)
	}

	type ResCandle struct {
		Open   float64 `json:"open"`
		Close  float64 `json:"close"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Volume int64   `json:"volume"`
	}
	res := make(map[string]ResCandle)

	for _, candle := range candles.GetCandles() {
		res[candle.Time.AsTime().String()] = ResCandle{
			Open:   helpers.ConvertQuotation(candle.Open),
			Close:  helpers.ConvertQuotation(candle.Close),
			High:   helpers.ConvertQuotation(candle.High),
			Low:    helpers.ConvertQuotation(candle.Low),
			Volume: candle.Volume,
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("%s : when convert to json", err)
	}
	fileName := fmt.Sprintf("candles.parsed.%s.%d.json", req.Figi, time.Now().Unix())
	os.WriteFile(fileName, data, os.ModePerm)
}
