package market_data

import (
	"net/http"

	"github.com/Ch2r1Ie/Stock-Bubble/app"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	market marketSrv
}

func NewHandler(market marketSrv) *Handler {
	return &Handler{market: market}
}

type StockInfoRequest struct {
	Symbol    []string `json:"symbol"`
	DateRange string   `json:"dateRange"`
	Interval  string   `json:"interval"`
}

func (handler *Handler) StockInfo(c *gin.Context) {
	var request StockInfoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Code: app.Err_BussinessErrors_2.Code, Message: app.Err_BussinessErrors_2.Message})
		return
	}

	response, err := handler.market.stocks(request.Symbol, request.DateRange, request.Interval)
	if err != nil {
		switch err {
		case &app.Err_UnExpected_StatusCode:
			c.JSON(http.StatusInternalServerError, app.Response{Code: app.Err_UnExpected_StatusCode.Code, Message: app.Err_UnExpected_StatusCode.Message})
			return
		case &app.Err_UnExpected_Response:
			c.JSON(http.StatusInternalServerError, app.Response{Code: app.Err_UnExpected_Response.Code, Message: app.Err_UnExpected_Response.Message})
			return
		default:
			c.JSON(http.StatusInternalServerError, app.Response{Code: app.Err_Unknown.Code, Message: app.Err_Unknown.Message})
			return
		}
	}

	c.JSON(http.StatusOK, app.Response{Code: app.Success.Code, Message: app.Success.Message, Data: response})
}
