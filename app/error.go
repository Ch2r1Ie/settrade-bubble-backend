package app

import (
	"encoding/json"
	"fmt"
)

var (
	Success = AppError{Code: 0000, Message: "Success"}
)

var (
	Err_Technical             = AppError{Code: 2000, Message: "Please contact tech support"}
	Err_UnExpected_StatusCode = AppError{Code: 2001, Message: "Unexpected status code from external api"}
	Err_UnExpected_Response   = AppError{Code: 2002, Message: "Unexpected response from external api"}

	Err_BussinessErrors_1 = AppError{Code: 1001, Message: "Not Found"}
	Err_BussinessErrors_2 = AppError{Code: 1004, Message: "Invalid data"}

	Err_Unknown = AppError{Code: 9999, Message: "Please contact admin"}
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("app error:%d|%s", e.Code, e.Message)
}

func (e *AppError) String() string {
	str, _ := json.Marshal(e)
	return string(str)
}

type ErrorResponse struct {
	Result ErrorResult `json:"result"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	var code int
	var message string
	switch v := err.(type) {
	case *AppError:
		code = v.Code
		message = v.Message
	default:
		code = Err_Unknown.Code
		message = Err_Unknown.Message
	}

	return ErrorResponse{
		Result: ErrorResult{
			Code:    code,
			Message: message,
		},
	}
}
