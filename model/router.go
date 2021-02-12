package model

import "net/http"

type IRouter interface {
	Route() http.Handler
}
