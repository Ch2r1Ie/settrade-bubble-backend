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
