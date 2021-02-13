package service

import (
	"context"
	"errors"
	"polybay/internal/apps/auth/service/db/auth_info"
	"polybay/internal/apps/auth/service/model"
	proto "polybay/internal/apps/auth/service/proto"
)

type handler struct {
	db model.IAuthInfo
}

func New(database model.Database) (proto.AuthServiceServer, error) {
	conn, err := auth_info.New(database)
	if err != nil {
		return nil, err
	}

	ret := &handler{db: conn}
	if err = ret.db.CreateTableIfNotExists(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (h *handler) Register(context.Context, *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, errors.New("not implemented")
}
