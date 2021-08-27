package utils

type (
	SuccessResponse struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		Status  int         `json:"status"`
		Meta    interface{} `json:"meta"`
	}
	ErrorResponse struct {
		Data      interface{} `json:"data"`
		Message   string      `json:"message"`
		Status    int         `json:"status"`
		ErrorCode string      `json:"errorCode"`
		Meta      interface{} `json:"meta"`
	}
)
