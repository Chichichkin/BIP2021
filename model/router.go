package model

import "net/http"

type IApp interface {
	Name() string
	Route() http.Handler
}

type IMiddleware interface {
	Handler(next http.Handler) http.Handler
}
