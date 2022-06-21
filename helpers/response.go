package helpers

type response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(msg string, data interface{}) *response {
	return &response{
		Error:   false,
		Message: msg,
		Data:    data,
	}
}

func ResponseError(msg string, data interface{}) *response {
	return &response{
		Error:   true,
		Message: msg,
		Data:    data,
	}
}
