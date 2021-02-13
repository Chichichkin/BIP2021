package auth

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"polybay/internal/apps/auth/service"
	proto "polybay/internal/apps/auth/service/proto"
	"polybay/model"
)

type auth struct {
	mux *runtime.ServeMux
}

func New() (model.IApp, error) {
	ret := &auth{
		mux: runtime.NewServeMux(),
	}

	if err := proto.RegisterAuthServiceHandlerServer(context.Background(), ret.mux, service.New()); err != nil {
		return nil, err
	}

	return ret, nil
}

func (a *auth) Route() http.Handler {
	return a.mux
}

func (a *auth) Name() string {
	return "/auth"
}
