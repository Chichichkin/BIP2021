package auth

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"polybay/internal/apps/auth/service"
	service_model "polybay/internal/apps/auth/service/model"
	proto "polybay/internal/apps/auth/service/proto"
	"polybay/model"
)

type auth struct {
	mux *runtime.ServeMux
}

func New(database model.Database) (model.IApp, error) {
	ret := &auth{
		mux: runtime.NewServeMux(),
	}

	micro, err := service.New(service_model.Database{
		Ip:       database.Ip,
		Port:     database.Port,
		Name:     database.Name,
		User:     database.User,
		Password: database.Password,
	})
	if err != nil {
		return nil, err
	}

	if err = proto.RegisterAuthServiceHandlerServer(context.Background(), ret.mux, micro); err != nil {
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
