package utils

type defaultResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

func SuccessResponse(data interface{}) *defaultResponse {
	response := defaultResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
	}

	return &response
}

func SuccessPaginationResponse(data, meta interface{}) *defaultResponse {
	response := defaultResponse{
		Code:    201,
		Message: "Success",
		Data:    data,
		Meta:    meta,
	}

	return &response
}

func ErrorProcessingDataResponse(message string) *defaultResponse {

	if message == "" {
		message = "Server error occurred."
	}

	response := defaultResponse{
		Code:    400,
		Message: message,
		Data:    nil,
	}

	return &response
}

func ErrorPayloadResponse() *defaultResponse {
	response := defaultResponse{
		Code:    401,
		Message: "payload error",
		Data:    nil,
	}

	return &response
}

func ErrorValidationResponse(message string) *defaultResponse {
	response := defaultResponse{
		Code:    500,
		Message: message,
		Data:    nil,
	}

	return &response
}

func ErrorDataNotFoundResponse() *defaultResponse {
	response := defaultResponse{
		Code:    401,
		Message: "Data not found",
		Data:    nil,
	}

	return &response
}
