package yahoo_finance

type Stock struct {
	Meta       Meta
	Timestamp  []int64
	Indicators Indicators
}

type Unadjclose struct {
	Unadjclose []float64
}

type Adjclose struct {
	Adjclose []float64
}

type Indicators struct {
	Quote      []Quote
	Unadjclose []Unadjclose
	Adjclose   []Adjclose
}

type Quote struct {
	Low    []float64
	Volume []float64
	High   []float64
	Open   []float64
	Close  []float64
}

type Meta struct {
	Currency             string
	Symbol               string
	ExchangeName         string
	InstrumentType       string
	FirstTradeDate       int64
	GmtOffset            int64 `json:"gmtoffset"`
	Timezone             string
	ExchangeTimezoneName string
	CurrentTradingPeriod CurrentTradingPeriod
	DataGranularity      string
	ValidRanges          []string
}

type CurrentTradingPeriod struct {
	Pre     TradingPeriod
	Regular TradingPeriod
	Post    TradingPeriod
}

type TradingPeriod struct {
	Timezone  string
	End       int64
	Start     int64
	GmtOffset int64 `json:"gmtoffset"`
}

type Info struct {
	Symbol           string `json:"symbol"`
	Currency         string `json:"currency"`
	ExchangeName     string `json:"exchangeName"`
	ExchangeTimezone string `json:"exchangeTimezone"`
	InstrumentType   string `json:"instrumentType"`
	FirstTradeDate   string `json:"firstTradeDate"`
	Price            Price  `json:"price"`
}

type Price struct {
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Volume           float64 `json:"volume"`
	Open             float64 `json:"open"`
	Close            float64 `json:"close"`
	PercentageChange float64 `json:"percentageChange"`
	Timestamp        string  `json:"timestamp"`
}
