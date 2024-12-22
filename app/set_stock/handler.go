package set_stock

import (
	tv "github.com/artlevitan/go-tradingview-ta"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	tradingview ITradingview
}

func NewHandler(tradingview ITradingview) *Handler {
	return &Handler{tradingview: tradingview}
}

func (handler *Handler) Transaction(c *gin.Context) {
	handler.tradingview.http(tv.Interval1Day)
}
