# Use the official Golang image as a build stage
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o bible-server .

# Use a minimal image for the final stage
FROM scratch

# Copy the compiled Go binary and necessary files from the builder stage
COPY --from=builder /app/bible-server /bible-server
COPY --from=builder /app/templates /templates
COPY --from=builder /app/static /static
COPY --from=builder /app/bible.db /bible.db

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/bible-server"]