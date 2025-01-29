package market_data

import (
	"github.com/Ch2r1Ie/Stock-Bubble/yahoo_finance.go"
)

var _ marketSrv = (*marketService)(nil)

type marketService struct {
	stock yahoo_finance.API
}

func NewMarketService(stock yahoo_finance.API) *marketService {
	return &marketService{stock: stock}
}

func (s *marketService) stocks(symbols []string, dateRange, interval string) ([]stock_info_response, error) {

	var stock_infos []stock_info_response
	for _, symbol := range symbols {

		stock_info, err := s.stock.Info(stock_symbols_input[symbol], dateRange, interval)
		if err != nil {
			return stock_infos, err
		}

		stock_infos = append(stock_infos, format(stock_info))

	}

	return stock_infos, nil
}
