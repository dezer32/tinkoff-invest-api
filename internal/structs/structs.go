package structs

import (
	"time"
)

type SharesResponse struct {
	Instruments []Share `json:"instruments"`
}

type Share struct {
	Figi                 string     `json:"figi"`
	Ticker               string     `json:"ticker"`
	ClassCode            string     `json:"classCode"`
	Isin                 string     `json:"isin"`
	Lot                  int        `json:"lot"`
	Currency             string     `json:"currency"`
	RatioLong            Quotation  `json:"klong"`
	RatioShort           Quotation  `json:"kshort"`
	RateLong             Quotation  `json:"dlong"`
	RateShort            Quotation  `json:"dshort"`
	RateLongMin          Quotation  `json:"dlongMin"`
	RateShortMin         Quotation  `json:"dshortMin"`
	HasShortEnabled      bool       `json:"shortEnabledFlag"`
	Name                 string     `json:"name"`
	Exchange             string     `json:"exchange"`
	IpoDate              time.Time  `json:"ipoDate"`
	IssueSize            string     `json:"issueSize"`
	RiskCountry          string     `json:"countryOfRisk"`
	RiskCountryName      string     `json:"countryOfRiskName"`
	Sector               string     `json:"sector"`
	IssueSizePlan        string     `json:"issueSizePlan"`
	Nominal              MoneyValue `json:"nominal"`
	TradingStatus        string     `json:"tradingStatus"`
	HasOtc               bool       `json:"otcFlag"`
	HasBuyAvailable      bool       `json:"buyAvailableFlag"`
	HasSellAvailable     bool       `json:"sellAvailableFlag"`
	HasDividendYield     bool       `json:"divYieldFlag"`
	ShareType            string     `json:"shareType"`
	MinPriceIncrement    Quotation  `json:"minPriceIncrement"`
	HasApiTradeAvailable bool       `json:"apiTradeAvailableFlag"`
	Uid                  string     `json:"uid"`
	RealExchange         string     `json:"realExchange"`
	PositionUid          string     `json:"positionUid"`
	HasIisAvailable      bool       `json:"forIisFlag"`
	IsOnlyQualified      bool       `json:"forQualInvestorFlag"`
	HasWeekend           bool       `json:"weekendFlag"`
	IsBlockedTCA         bool       `json:"blockedTcaFlag"`
	First1DayCandleDate  time.Time  `json:"first1dayCandleDate"`
	First1MinCandleDate  time.Time  `json:"first1minCandleDate"`
}

type MoneyValue struct {
	Nano     int    `json:"nano"`
	Currency string `json:"currency"`
	Units    string `json:"units"`
}

type Quotation struct {
	Nano  int    `json:"nano"`
	Units string `json:"units"`
}
