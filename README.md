# Polybay
Main repository for our project

## Protoc
Для начала требуется накатить protoc https://github.com/protocolbuffers/protobuf/releases/tag/3.15.0-rc1

После накатки на свою прекрасную шинду добавляем в переменные окружения.

Далее идем в папку вашего компилятора go (например go1.15), и заходим в папочку bin.

Открываем консоль и пишем:

```
go get google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

Чтоб сгенерить протобаф выполняем bat скрипт.