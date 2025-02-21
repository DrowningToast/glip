package common

type HTTPResponse struct {
	Result interface{} `json:"result,omitempty"`

	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type PaginatedResult struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

var EmptyHTTPResponse = HTTPResponse{
	Result: struct{}{},
}
