package rest

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}

func GetSuccess(statusCode int, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       data,
		Message:    "Success",
		Success:    true,
	}
}

func GetError(statusCode int, msg string) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    msg,
	}
}
