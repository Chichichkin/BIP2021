package model

import "net/http"

type IApp interface {
	IRouter
	Name() string
}

type IRouter interface {
	Route() http.Handler
}
