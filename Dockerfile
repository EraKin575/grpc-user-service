FROM golang:1.22-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc-user-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/grpc-user-service .
RUN chmod +x /root/grpc-user-service
EXPOSE 50051
CMD ["/root/grpc-user-service"]