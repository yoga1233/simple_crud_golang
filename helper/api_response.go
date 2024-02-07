package helper

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta        `json:"metadata"`
	Data interface{} `json:"response"`
}
type ResponseFailure struct {
	Meta Meta `json:"metadata"`
}

func ApiResponseSuccess(message string, code int, data interface{}) Response {
	meta := Meta{Code: code, Message: message}

	jsonResponse := Response{Meta: meta, Data: data}

	return jsonResponse
}
func ApiResponseFailure(message string, code int) ResponseFailure {
	meta := Meta{Message: message, Code: code}

	return ResponseFailure{Meta: meta}
}
