package helper

type HttpErrorResponse struct {
	Errors    []interface{}
	Timestamp string
}

type HttpError struct {
	error
	Message []string
}

func New(err error, status int, httpMessage []string) *HttpError {
	return &HttpError{
		error:   err,
		Message: httpMessage,
	}
}
