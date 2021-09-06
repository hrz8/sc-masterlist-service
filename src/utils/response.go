package utils

type (
	// SuccessResponse represents structure of success rest response
	SuccessResponse struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		Status  int         `json:"status"`
		Meta    interface{} `json:"meta"`
	}

	// ErrorResponse represents structure of failed rest response
	ErrorResponse struct {
		Data      interface{} `json:"data"`
		Message   string      `json:"message"`
		Status    int         `json:"status"`
		ErrorCode string      `json:"errorCode"`
		Meta      interface{} `json:"meta"`
	}

	// ListMetaResponse represents structure of meta rest response
	ListMetaResponse struct {
		Count int   `json:"count"`
		Total int64 `json:"total"`
	}
)
