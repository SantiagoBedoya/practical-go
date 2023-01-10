DB_URI := "postgres://postgres:root@localhost:5432/auth?sslmode=disable"

generate:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. internal/pb/*.proto

grpcui:
	grpcui -plaintext localhost:5001

migrate:
	migrate -database "${DB_URI}" -path db/migrations up

migrate-rollback:
	migrate -database "${DB_URI}" -path db/migrations down

keys:
	openssl genrsa -out rsa-keys/id_rsa 4096
	openssl rsa -in rsa-keys/id_rsa -pubout -out rsa-keys/id_rsa.pub

start:
	go run cmd/main.go -config .config.cfg