# Use the official Golang image as the base image
FROM golang:1.24.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Install Air for live-reloading
RUN go install github.com/air-verse/air@latest

# Copy the Go module files
COPY go.mod .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o inventory-service .

# Expose the port your app runs on (if applicable)
EXPOSE 8081

# Command to run the application
CMD ["./inventory-service"]
