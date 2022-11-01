package error_code

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"strings"
)

type ErrorCode struct {
	Status       int32  `json:"status,omitempty"`
	ErrorCode    int32  `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	Exception    string `json:"exception,omitempty"`
	WaitFor      int    `json:"wait_for,omitempty"`
}

type SuccessCode struct {
	Status         int32  `json:"status,omitempty"`
	SuccessCode    int32  `json:"success_code,omitempty"`
	SuccessMessage string `json:"success_message,omitempty"`
	WaitFor        int    `json:"wait_for,omitempty"`
}

func NewError(code int32, message string, line string) (err *ErrorCode) {
	httpCode := int32(code) / 1000
	if httpCode >= 1000 {
		httpCode = httpCode / 10
	}

	message = strings.TrimSpace(message)
	if name, ok := ERROR_CODE_NAME[int32(code)]; ok {
		if message == "" {
			message = name
		}

		if httpCode <= 500 {
			err = &ErrorCode{
				Status:       httpCode,
				ErrorCode:    code,
				ErrorMessage: fmt.Sprintf("%s: code = %d, message = %s", name, code, message),
				WaitFor:      10,
				Exception:    fmt.Sprintf("Time: %s, Line: %s", base.TimeCurrent(), line),
			}
		} else {
			err = &ErrorCode{
				Status:       httpCode,
				ErrorCode:    code,
				ErrorMessage: fmt.Sprintf("%s: code = %d, message = %s", name, code, message),
				WaitFor:      30,
				Exception:    fmt.Sprintf("Time: %s, Line: %s", base.TimeCurrent(), line),
			}
		}
	} else {
		code = int32(ERROR_INTERNAL)
		httpCode = int32(code) / 1000
		if httpCode >= 1000 {
			httpCode = httpCode / 10
		}

		err = &ErrorCode{
			Status:       httpCode,
			ErrorCode:    code,
			ErrorMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: code = %d, message = %s", code, message),
			WaitFor:      60,
			Exception:    fmt.Sprintf("Time: %s, Line: %s", base.TimeCurrent(), line),
		}
	}

	return
}

func NewSuccess(code int32, message string, waitFor int) (err *SuccessCode) {
	httpCode := int32(code) / 1000
	if httpCode >= 1000 {
		httpCode = httpCode / 10
	}

	if name, ok := ERROR_CODE_NAME[int32(code)]; ok {

		if message == "" {
			message = name
		}

		err = &SuccessCode{
			Status:         httpCode,
			SuccessCode:    code,
			SuccessMessage: fmt.Sprintf("%s: code = %d, message = %s", name, code, message),
			WaitFor:        waitFor,
		}
	} else {
		code = int32(ERROR_CODE_OK)
		httpCode = int32(code) / 1000
		if httpCode >= 1000 {
			httpCode = httpCode / 10
		}

		err = &SuccessCode{
			Status:         httpCode,
			SuccessCode:    code,
			SuccessMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: code = %d, message = %s", code, message),
			WaitFor:        waitFor,
		}
	}

	return
}
