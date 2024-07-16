FROM golang:1.22-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o grpc-user-service ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/grpc-user-service .

EXPOSE 50051

CMD ["./grpc-user-service"]
