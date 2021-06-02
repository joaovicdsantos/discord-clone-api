package exception

// HttpError class of any error that may occur in the request
type HttpError struct {
	Err        error
	StatusCode int
}
