package market_data

import "time"

type SparklineData struct {
	Price []float64 `json:"price"`
}

type ROI struct {
	Times    int    `json:"times"`
	Currency string `json:"currency"`

	Percentage float64 `json:"percentage"`
}

type stock_info_response struct {
	ID                                 string        `json:"id"`
	Symbol                             string        `json:"symbol"`
	Name                               string        `json:"name"`
	Image                              string        `json:"image"`
	CurrentPrice                       float64       `json:"current_price"`
	MarketCap                          float64       `json:"market_cap"`
	MarketCapRank                      int           `json:"market_cap_rank"`
	FullyDilutedValuation              float64       `json:"fully_diluted_valuation"`
	TotalVolume                        float64       `json:"total_volume"`
	High24h                            float64       `json:"high_24h"`
	Low24h                             float64       `json:"low_24h"`
	PriceChange24h                     float64       `json:"price_change_24h"`
	PriceChangePercentage24h           float64       `json:"price_change_percentage_24h"`
	MarketCapChange24h                 float64       `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h       float64       `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                  float64       `json:"circulating_supply"`
	TotalSupply                        float64       `json:"total_supply"`
	MaxSupply                          *float64      `json:"max_supply,omitempty"`
	Ath                                float64       `json:"ath"`
	AthChangePercentage                float64       `json:"ath_change_percentage"`
	AthDate                            time.Time     `json:"ath_date"`
	Atl                                float64       `json:"atl"`
	AtlChangePercentage                float64       `json:"atl_change_percentage"`
	AtlDate                            time.Time     `json:"atl_date"`
	Roi                                *ROI          `json:"roi,omitempty"`
	LastUpdated                        time.Time     `json:"last_updated"`
	SparklineIn7d                      SparklineData `json:"sparkline_in_7d"`
	PriceChangePercentage1hInCurrency  float64       `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage1yInCurrency  float64       `json:"price_change_percentage_1y_in_currency"`
	PriceChangePercentage24hInCurrency float64       `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage30dInCurrency float64       `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage7dInCurrency  float64       `json:"price_change_percentage_7d_in_currency"`
}
