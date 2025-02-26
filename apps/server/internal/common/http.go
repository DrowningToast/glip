package common

type HTTPResponse struct {
	Result interface{} `json:"result,omitempty"`

	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type PaginatedResult[T interface{}] struct {
	Count int `json:"count"`
	Items []T `json:"items"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

var EmptyHTTPResponse = HTTPResponse{
	Result: struct{}{},
}
