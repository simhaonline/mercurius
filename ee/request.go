package ee

import (
	"fmt"
	"net/http"
)

type ErrKeyNotFound struct {
	Cause error
	Key   string
}

func (e ErrKeyNotFound) Error() string {
	return fmt.Sprintf("key '%s' not found", e.Key)
}

func (e ErrKeyNotFound) Unwrap() error {
	return e.Cause
}

type Request interface {
	// Params evaluates in an implementation specific order multiple sources of parameters
	// and returns the "first" occurrence of the key, whatever "first" means in
	// the concrete implementation. If the key is not found, an ErrKeyNotFound is returned.
	Param(key string) (interface{}, error)
}

type httpRequestImpl struct {
	request *http.Request
}

// Param evaluates Path Para
func (h httpRequestImpl) Param(key string) (interface{}, error) {
	return nil, nil
}

func NewRequest(request *http.Request) Request {
	return httpRequestImpl{request}
}
