FROM golang:1.18-alpine AS BUILDER
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest

WORKDIR /root
COPY --from=BUILDER /app/server .

CMD [ "./server" ]