package errs

type Err struct {
	// HTTP status code
	StatusCode int
	// Error code (eg. "DUPLICATE")
	Code string
	// Any string
	Message string
}

func (e Err) Error() string {
	return e.Message
}

var _ error = (*Err)(nil)

var (
	ErrInternal = Err{
		StatusCode: 500,
		Code:       "INTERNAL",
		Message:    "An internal server error occurred",
	}
	ErrInvalidArgument = Err{
		StatusCode: 400,
		Code:       "INVALID_ARGUMENT",
		Message:    "Invalid argument",
	}
	ErrNotFound = Err{
		StatusCode: 404,
		Code:       "NOT_FOUND",
		Message:    "Resource not found",
	}
	ErrUnauthorized = Err{
		StatusCode: 401,
		Code:       "UNAUTHORIZED",
		Message:    "Unauthorized",
	}
	ErrInvalidQueryString = Err{
		StatusCode: 400,
		Code:       "INVALID_QUERY_STRING",
		Message:    "Invalid query string",
	}
	ErrForbidden = Err{
		StatusCode: 403,
		Code:       "FORBIDDEN",
		Message:    "Forbidden",
	}

	ErrDuplicate = Err{
		StatusCode: 400,
		Code:       "DUPLICATE",
		Message:    "Duplicate resource",
	}
	ErrInvalidBody = Err{
		StatusCode: 400,
		Code:       "INVALID_BODY",
		Message:    "Invalid body",
	}
)
