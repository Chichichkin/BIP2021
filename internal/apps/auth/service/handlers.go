package service

import (
	"context"
	"errors"
	proto "polybay/internal/apps/auth/service/proto"
)

type handler struct {
}

func New() proto.AuthServiceServer {
	return &handler{}
}

func (h *handler) Register(context.Context, *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	return nil, errors.New("not implemented")
}
