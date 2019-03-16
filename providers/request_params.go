package providers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RequestParamsInterface ...
type RequestParamsInterface interface {
	Get(req *http.Request) map[string]string
}

// RequestParamsProvider ...
type RequestParamsProvider struct{}

// Get ...
func (r *RequestParamsProvider) Get(req *http.Request) map[string]string {
	return mux.Vars(req)
}