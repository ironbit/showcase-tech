# Select parent image (golang) for building the application
FROM golang:alpine AS builder

# Change to working directory /build
WORKDIR /build/

# Copy and download dependency using go mod
COPY app .
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-bin main.go


# Select parent image (debian) for running the application
FROM alpine:latest

# Change to working directory /deploy
WORKDIR /deploy/

# Copy application binary
COPY --from=builder /build/app-bin .

# Copy config file
RUN mkdir /deploy/conf
COPY --from=builder /build/conf/app.conf /deploy/conf

# Open the port in the container
EXPOSE $APP_PORT

# Execute binary
CMD ["nohup", "/deploy/app-bin"]
