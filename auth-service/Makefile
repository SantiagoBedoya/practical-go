PORT := 8080
DB_URI := postgres://postgres:root@localhost:5432/accounts?sslmode=disable
JWT_SECRET := Sup3rS3cr37

start:
	PORT=${PORT} \
	DB_URI=${DB_URI} \
	JWT_SECRET=${JWT_SECRET} \
	go run main.go

migrate:
	migrate -database "${DB_URI}" -path db/migrations up

migrate-rollback:
	migrate -database "${DB_URI}" -path db/migrations down

migrate-fix:
	migrate -database "${DB_URI}" -path db/migrations force 1