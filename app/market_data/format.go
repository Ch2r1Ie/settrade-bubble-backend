package market_data

import (
	"math"
	"time"

	"github.com/Ch2r1Ie/Stock-Bubble/yahoo_finance.go"
)

const (
	time_format = "15:04:05 02/01/2006"
)

func format_response(stock yahoo_finance.Stock) yahoo_finance.Info {
	return yahoo_finance.Info{
		Currency:         stock.Meta.Currency,
		Symbol:           stock.Meta.Symbol,
		InstrumentType:   stock.Meta.InstrumentType,
		ExchangeName:     stock.Meta.ExchangeName,
		ExchangeTimezone: stock.Meta.ExchangeTimezoneName,
		FirstTradeDate:   formatTime(stock.Meta.FirstTradeDate, time_format),
		Price: yahoo_finance.Price{
			High:             format_decimal(stock.Indicators.Quote[0].High[lastIndex(len(stock.Indicators.Quote[0].High))]),
			Low:              format_decimal(stock.Indicators.Quote[0].Low[lastIndex(len(stock.Indicators.Quote[0].Low))]),
			Volume:           format_decimal(stock.Indicators.Quote[0].Volume[lastIndex(len(stock.Indicators.Quote[0].Volume))]),
			Open:             format_decimal(stock.Indicators.Quote[0].Open[lastIndex(len(stock.Indicators.Quote[0].Open))]),
			Close:            format_decimal(stock.Indicators.Quote[0].Close[lastIndex(len(stock.Indicators.Quote[0].Close))]),
			PercentageChange: format_decimal(calculatePercentage(stock)),
			Timestamp:        formatTime(stock.Timestamp[lastIndex(len(stock.Timestamp))], time_format),
		},
	}
}

func formatTime(timestamp int64, template string) string {
	return time.Unix(timestamp, 0).Format(template)
}

func format_decimal(number float64) float64 {
	return math.Round(number*100) / 100
}

func calculatePercentage(stock yahoo_finance.Stock) float64 {

	if len(stock.Indicators.Quote) == 0 {
		return 0
	}

	closePrices := stock.Indicators.Quote[0].Close
	if len(closePrices) < 2 {
		return 0
	}

	oldPrice := stock.Indicators.Quote[0].Open[lastIndex(len(stock.Indicators.Quote[0].Open))]
	newPrice := stock.Indicators.Quote[0].Close[lastIndex(len(stock.Indicators.Quote[0].Close))]

	if oldPrice == 0 {
		return 0
	}

	percentageChange := ((newPrice - oldPrice) / oldPrice) * 100
	return percentageChange
}

func lastIndex(maxLength int) int {
	return maxLength - 1
}
