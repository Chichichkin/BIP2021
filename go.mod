module polybay

go 1.15

require (
	github.com/Chichichkin/authentication-service v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v1.5.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/hashicorp/consul/api v1.8.1
	github.com/lib/pq v1.9.0 // indirect
	google.golang.org/genproto v0.0.0-20210212180131-e7f2df4ecc2d // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/Chichichkin/authentication-service => ../authentication-service
