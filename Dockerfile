# Use the official Golang image as the base image
FROM golang:1.21-alpine3.19 AS builder

# Set the working directory inside the container
WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o bmkg_api .

# Start a new stage from scratch
FROM alpine:latest
# Set the Current Working Directory inside the container
WORKDIR /app

COPY --from=builder /app/bmkg_api .
COPY .env .

EXPOSE 8000

# Run the application
CMD ["./bmkg_api"]