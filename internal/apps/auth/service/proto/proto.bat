set GOPATH=C:\Users\dds0l\go
set files=auth.proto

protoc -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis^
  --go_out=plugins=grpc:./ ./*.proto^
  --grpc-gateway_out=logtostderr=true:./^
  %auth.proto%

protoc  -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis^
    -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway^
    --swagger_out=allow_merge=true,merge_file_name=api:./^
    --go_out=plugins=grpc:./ ./*.proto^
    %files%