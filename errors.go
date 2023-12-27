package http_client_go

import "errors"

var (
	ErrMethodNotSupport = errors.New("method has not been supported")
	ErrHostIsEmpty      = errors.New("host is empty")
	ErrURLIsEmpty       = errors.New("url can not be empty")
)
