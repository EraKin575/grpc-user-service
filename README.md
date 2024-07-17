# gRPC User Service

This is a gRPC service for managing user details, including a search capability based on specific criteria.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Service](#running-the-service)
- [Testing the Service](#testing-the-service)
- [Using Docker Compose](#using-docker-compose)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Features

- Fetch user details based on user ID.
- Retrieve a list of user details based on a list of user IDs.
- Search user details based on specific criteria (city, phone number, marital status).

## Prerequisites

- Docker
- Docker Compose
- `grpcurl` (optional, for testing the gRPC service)

## Installation

1. Clone the repository:
    ```bash
    git clone <your-repository-url>
    cd grpc-user-service
    ```

## Running the Service

### Using Docker Compose


1. Build and run the services with Docker Compose:
    ```bash
    docker-compose up --build
    ```

This will build the Docker image and start the gRPC service on port 50051.

## Testing the Service

You can test the gRPC service using `grpcurl`:

1. **Install `grpcurl`**:

    - For Linux:
      ```bash
      sudo apt install grpcurl
      ```

    - For macOS (using Homebrew):
      ```bash
      brew install grpcurl
      ```

2. **Interact with the gRPC service**:

    - Get user by ID:
      ```bash
      grpcurl -plaintext -d '{"id": 1}' localhost:50051 user.UserService/GetUser
      ```

    - List users by IDs:
      ```bash
      grpcurl -plaintext -d '{"ids": [1, 2]}' localhost:50051 user.UserService/ListUsers
      ```

    - Search users by criteria:
      ```bash
      grpcurl -plaintext -d '{"city": "LA"}' localhost:50051 user.UserService/SearchUsers
      ```

## API Endpoints

### GetUser

Fetch user details based on user ID.

- **Request**: `{"id": int}`
- **Response**: `{"user": {"id": int, "fname": string, "city": string, "phone": int64, "height": float32, "married": bool}}`

### ListUsers

Retrieve a list of user details based on a list of user IDs.

- **Request**: `{"ids": [int]}`
- **Response**: `{"users": [{"id": int, "fname": string, "city": string, "phone": int64, "height": float32, "married": bool}]}`

### SearchUsers

Search user details based on specific criteria (city, phone number, marital status).

- **Request**: `{"city": string, "phone": int64, "married": bool}`
- **Response**: `{"users": [{"id": int, "fname": string, "city": string, "phone": int64, "height": float32, "married": bool}]}`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
