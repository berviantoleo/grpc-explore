# GRPC Explore

## Tools

- .NET 8.0 SDK
- Go
- Docker
- Docker Compose

## Run the GRPC Server

### Using Docker & Docker Compose

- Run `docker compose server up -d`

### Using .NET SDK

- Run `dotnet run --project GRPCExample`

## Run the .NET GRPC Client

Ensure you've set the `SERVER_URL`.

### Using Docker & Docker Compose

- Run `docker compose dotnetclient up`

### Using .NET SDK

- Run `dotnet run --project GRPC.Client`

## Run the Go GRPC Client

Ensure you've set the `SERVER_URL`.

### Using Docker & Docker Compose

- Run `docker compose goclient up`

### Using Go SDK

- Run `go run main.go`

## LICENCE

MIT