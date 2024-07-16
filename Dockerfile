# Stage 1: Build the Go binary
FROM golang:1.16-alpine AS builder

# Install necessary packages
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o grpc-user-service ./main.go

# Stage 2: Create a small image with only the necessary binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/grpc-user-service .

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./grpc-user-service"]
