# Select parent image (golang) for building the application
FROM golang:alpine AS builder

# Change to working directory /build
WORKDIR /build/

# Copy and download dependency using go mod
COPY app .
RUN go mod download

# Build the application
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app-bin main.go


# Select parent image (alpine) for running the application
FROM scratch as bin

# Change to working directory /deploy
WORKDIR /deploy/

# Copy application binary
COPY --from=builder /build/app-bin .

# Open the port in the container
EXPOSE $APP_PORT

# Execute binary
CMD ["./app-bin"]  
