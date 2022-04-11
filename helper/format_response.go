package helper

type FormatResponse struct {
	Meta MetaData    `json:"meta"`
	Data interface{} `json:"data"`
}

type MetaData struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseFormat(message string, code int, status string, data interface{}) FormatResponse {
	meta := MetaData{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := FormatResponse{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
