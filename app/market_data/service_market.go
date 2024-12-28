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

func (s *marketService) stocks(symbols []string, dateRange, interval string) (*[]yahoo_finance.Info, error) {

	var stock_infos []yahoo_finance.Info
	for _, symbol := range symbols {

		if symbol != "" {
			stock_info, err := s.stock.Info(stock_symbols_input[symbol], dateRange, interval)
			if err != nil {
				return nil, err
			}

			stock_infos = append(stock_infos, format_response(stock_info))
		}

	}

	return &stock_infos, nil
}
