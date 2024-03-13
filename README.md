# gRPC serive in Go

Install protobuf using brew

```sh
brew install protobuf protoc-gen-go protoc-gen-go-grpc
```

Alternatively install the protoc plugins using golang

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Install gPRC package from google

```sh
go get google.golang.org/grpc
```

Generate the gRPC code

```sh
make generate
```

Run the server

```sh
go run server/main.go
```

Test with grpcui

Install

```sh
brew install grpcui
```

Test
If you're wanting to connect locally, you'll need to use the -plaintext flag with both grpcui and grpcurl
```sh
grpcui --plaintext 127.0.0.1:8080
```

Build and Run gRPC Server in Docker Container

```sh
docker build -t grpc-server .
docker run --name my_grpc_server -p 8080:8080 grpc-server
```