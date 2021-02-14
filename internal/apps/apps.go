package apps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"polybay/internal/apps/default"
	"polybay/internal/middlewares/authentification"
	"polybay/model"
)

type app struct {
	apps           []model.IApp
	authMiddleware model.IMiddleware
}

func New(config *model.Config) (*app, error) {
	ret := &app{apps: make([]model.IApp, 0, 100)}
	mid, err := authentification.New("localhost:5300")
	if err != nil {
		return nil, err
	}
	ret.authMiddleware = mid

	tmp, err := _default.New("/auth", "localhost:3000", nil)
	if err != nil {
		return nil, err
	}
	ret.apps = append(ret.apps, tmp)

	return ret, nil
}

func (a *app) Route() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	for _, app := range a.apps {
		r.Mount(app.Name(), app.Route())
	}
	return r
}
