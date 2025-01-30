package errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	error
	StatusCode int
}

func NewHttpError(error error, statusCode int) *HttpError {
	return &HttpError{error: error, StatusCode: statusCode}
}

var (
	CoinNotFoundError    = fmt.Errorf("coin not found")
	ChainNotFoundError   = fmt.Errorf("chain not found")
	UnexpectedError      = fmt.Errorf("unexpected error")
	TooManyRequestsError = fmt.Errorf("too many requests")
	NotImplementedError  = NewHttpError(fmt.Errorf("NotImplementedError"), http.StatusNotImplemented)
	BadRequestError      = NewHttpError(fmt.Errorf("BadRequestError"), http.StatusBadRequest)
	InternalServerError  = NewHttpError(fmt.Errorf("InternalServerError"), http.StatusInternalServerError)
	NotFoundError        = NewHttpError(fmt.Errorf("NotFoundError"), http.StatusNotFound)
)
