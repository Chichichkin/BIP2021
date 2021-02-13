package apps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"polybay/internal/apps/auth"
	"polybay/model"
)

type app struct {
	apps []model.IApp
}

func New() (model.IRouter, error) {
	apps := make([]model.IApp, 0, 100)
	tmp, err := auth.New()
	if err != nil {
		return nil, err
	}
	apps = append(apps, tmp)

	return &app{
		apps: apps,
	}, nil
}

func (a *app) Route() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	for _, app := range a.apps {
		return app.Route()
	}
	return r
}
