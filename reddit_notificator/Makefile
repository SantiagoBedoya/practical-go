generate:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./internal/pb/*.proto

start:
	go run cmd/main.go