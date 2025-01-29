package market_data

type marketSrv interface {
	stocks(symbols []string, dateRange, interval string) ([]stock_info_response, error)
}
