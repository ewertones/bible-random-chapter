# Use the official Golang image as the base image
FROM golang:1.24-alpine

# Install build dependencies
RUN apk update && apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o bible-server .

# Expose the port the app runs on
EXPOSE 8080

# # Command to run the executable
ENTRYPOINT ["/app/bible-server"]