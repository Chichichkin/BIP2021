package authentification

import (
	"context"
	pb "github.com/Chichichkin/authentication-service/service/proto"
	"log"
	"polybay/model"

	"google.golang.org/grpc"
	"net/http"
)

type auth struct {
	cli pb.AuthServiceClient
}

func New(address string) (model.IMiddleware, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ret := &auth{
		cli: pb.NewAuthServiceClient(conn),
	}

	return ret, nil
}

func (a *auth) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := a.cli.Register(context.Background(), &pb.RegisterRequest{
			Email:           "emial",
			Password:        "passw",
			RetypedPassword: "passw",
		})
		if err != nil {
			log.Println(err.Error())
		}
		next.ServeHTTP(w, r)
	})
}
