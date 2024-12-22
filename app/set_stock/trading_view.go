package set_stock

import (
	"fmt"

	tv "github.com/artlevitan/go-tradingview-ta"
)

type ITradingview interface {
	http(tradingview_interval string) (float64, float64, float64, error)
}

type tradingview struct{}

func NewTrade() *tradingview {
	return &tradingview{}
}

func (s *tradingview) http(tradingview_interval string) (float64, float64, float64, error) {

	var tv tv.TradingView
	if err := tv.Get(SYMBOL, tradingview_interval); err != nil {
		return 0, 0, 0, err
	}

	fmt.Println(tv.Value.Oscillators.AO.Prev1)
	fmt.Println(tv.Value.Oscillators.AO.Prev2)
	fmt.Println(tv.Value.Oscillators.AO.Value)
	fmt.Println(((tv.Value.Oscillators.AO.Value - tv.Value.Oscillators.AO.Prev1) / tv.Value.Oscillators.AO.Prev1) * 100)

	return tv.Value.Prices.High, tv.Value.Prices.Low, tv.Value.Prices.Close, nil
}
