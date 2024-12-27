package market_data

import (
	"net/http"

	"github.com/Ch2r1Ie/Stock-Bubble/app"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	yahoo yahooFinance
}

func NewHandler(
	yahoo yahooFinance,
) *Handler {
	return &Handler{
		yahoo: yahoo,
	}
}

type StockInfoRequest struct {
	Symbol    string `json:"symbol"`
	DateRange string `json:"dateRange"`
	Interval  string `json:"interval"`
}

func (handler *Handler) StockInfo(c *gin.Context) {
	var request StockInfoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	response, err := handler.yahoo.stock(request.Symbol, request.DateRange, request.Interval, nil)
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
