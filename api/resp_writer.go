package api

import (
	"github.com/Kephas73/lib-kephas/error_code"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ObjectExtend interface {
	String()
}

type Response struct {
	Message    string                  `json:"Message,omitempty"`
	Data       interface{}             `json:"Data,omitempty"`
	DataExtend map[string]ObjectExtend `json:"DataExtend,omitempty"`
	Status     int                     `json:"Status,omitempty"`
}

func WriteSuccess(c echo.Context, v interface{}) error {
	res := Response{
		Message: "Success!!",
		Data:    v,
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusOK, res)
}

func WriteSuccessEmptyContent(c echo.Context) error {
	res := Response{
		Message: "Success!!",
		Data:    nil,
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusOK, res)
}

func WriteSuccessWithExtend(c echo.Context, data interface{}, extend map[string]ObjectExtend) error {
	res := Response{
		Message:    "Success!!",
		Data:       data,
		DataExtend: extend,
		Status:     http.StatusOK,
	}

	return c.JSON(http.StatusOK, res)
}

func WriterSuccessOther(c echo.Context, success *error_code.SuccessCode) error {
	res := Response{
		Message: "Success!!",
		Data:    success,
		Status:  int(success.Status),
	}

	// Return
	return c.JSON(int(success.Status), res)
}

func WriterSuccessOtherWithExtend(c echo.Context, success *error_code.SuccessCode, extend map[string]ObjectExtend) error {
	res := Response{
		Message:    "Success!!",
		Data:       success,
		DataExtend: extend,
		Status:     int(success.Status),
	}

	// Return
	return c.JSON(int(success.Status), res)
}

func WriteError(c echo.Context, err *error_code.ErrorCode) error {
	res := Response{
		Message: "Failure!!",
		Data:    err,
		Status:  int(err.Status),
	}

	// Return
	return c.JSON(int(err.Status), res)
}

func WriteErrorWithExtend(c echo.Context, err *error_code.ErrorCode, extend map[string]ObjectExtend) error {
	res := Response{
		Message:    "Failure!!",
		Data:       err,
		DataExtend: extend,
		Status:     int(err.Status),
	}

	// Return
	return c.JSON(int(err.Status), res)
}
