.PHONY: generate

generate:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. unary/proto/*.proto
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. server_streaming/proto/*.proto
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. client_streaming/proto/*.proto
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. bidirectional_streaming/proto/*.proto