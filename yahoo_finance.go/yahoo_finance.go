package yahoo_finance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/Ch2r1Ie/Stock-Bubble/app"
)

type openAPI struct {
	url string
}

func NewOpenAPI(url string) *openAPI {
	return &openAPI{url: url}
}

var (
	DebugLogging = false
)

func (y *openAPI) Info(symbol string, dateRange, interval string) (Stock, error) {

	var stock Stock
	resp, err := exec(symbol, dateRange, interval, y.url)
	if err != nil {
		return stock, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return stock, &app.Err_UnExpected_StatusCode
	}

	var target struct {
		Chart struct {
			Result []Stock
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return stock, err
	}

	if len(target.Chart.Result) != 1 {
		return stock, &app.Err_UnExpected_Response
	}

	return target.Chart.Result[0], nil
}

func exec(symbol, dateRange, interval, url string) (*http.Response, error) {

	tmpl, err := template.New("YF-API").Parse(url)
	if err != nil {
		return nil, err
	}

	p := struct {
		Symbol   string
		Range    string
		Interval string
	}{
		Symbol:   symbol,
		Range:    dateRange,
		Interval: interval,
	}
	var result bytes.Buffer
	if err := tmpl.Execute(&result, p); err != nil {
		return nil, err
	}

	debug(result.String())
	return http.Get(result.String())
}

func debug(str string) {
	if !DebugLogging {
		return
	}

	fmt.Println(str)
}
