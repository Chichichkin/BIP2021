package _default

import (
	"github.com/go-chi/chi"
	"net/http"
	"polybay/internal/wrapper"
	"polybay/model"
	"polybay/tools"
)

type def struct {
	route      string
	remote     string
	middleware func(handler http.Handler) http.Handler
}

func New(route, remote string, middleware func(handler http.Handler) http.Handler) (model.IApp, error) {
	return &def{route: route, remote: remote, middleware: middleware}, nil
}

func (d *def) Route() http.Handler {
	r := chi.NewRouter()
	if d.middleware != nil {
		r.Use(d.middleware)
	}

	r.HandleFunc("/*", d.handler)
	return r
}

func (d *def) handler(w http.ResponseWriter, r *http.Request) {
	resp, err := tools.Resend(r, d.remote)
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}
	wrapper.SuccessResponse(w, resp)
}

func (d *def) Name() string {
	return d.route
}
