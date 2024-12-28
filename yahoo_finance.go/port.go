package yahoo_finance

//go:generate mockgen -source=port.go -destination=../app/market_data/mock/yahoo_finance_port_mock.go -package=mock

type API interface {
	Info(symbol string, dateRange, interval string) (Stock, error)
}
