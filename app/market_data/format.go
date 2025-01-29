package market_data

import (
	"strings"

	"github.com/Ch2r1Ie/Stock-Bubble/yahoo_finance.go"
)

func format(stock yahoo_finance.Stock) stock_info_response {

	return stock_info_response{
		ID:                                 modifiedBK(stock_symbols_output[stock.Meta.Symbol]),
		Symbol:                             stock_symbols_output[stock.Meta.Symbol],
		Name:                               modifiedBK(stock_symbols_output[stock.Meta.Symbol]),
		CurrentPrice:                       stock.Indicators.Quote[0].Close[lastIndex(len(stock.Indicators.Quote[0].Close))],
		PriceChange24h:                     calculatePriceChange(stock),
		PriceChangePercentage24h:           calculatePercentageChange(stock),
		PriceChangePercentage1hInCurrency:  calculatePercentageChange(stock),
		PriceChangePercentage7dInCurrency:  calculatePercentageChange(stock),
		PriceChangePercentage24hInCurrency: calculatePercentageChange(stock),
		PriceChangePercentage30dInCurrency: calculatePercentageChange(stock),
		PriceChangePercentage1yInCurrency:  calculatePercentageChange(stock),
	}
}

func calculatePercentageChange(stock yahoo_finance.Stock) float64 {

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

func calculatePriceChange(stock yahoo_finance.Stock) float64 {

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

	priceChange := newPrice - oldPrice
	return priceChange
}

func lastIndex(maxLength int) int {
	return maxLength - 1
}

func modifiedBK(original string) string {
	return strings.TrimSuffix(original, ".BK")
}
