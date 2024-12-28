package market_data

import "github.com/Ch2r1Ie/Stock-Bubble/yahoo_finance.go"

type marketSrv interface {
	stocks(symbols []string, dateRange, interval string) (*[]yahoo_finance.Info, error)
}
