package exception

type HttpError struct {
	Err        error
	StatusCode int
}
