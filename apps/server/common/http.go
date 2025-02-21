package common

type HTTPResponse struct {
	Result interface{} `json:"result,omitempty"`

	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

var EmptyHTTPResponse = HTTPResponse{
	Result: struct{}{},
}
