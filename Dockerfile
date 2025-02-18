# Use the official Golang image as the base image
FROM golang:1.22.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o e-commerce .

# Expose the port your app runs on (if applicable)
EXPOSE 8080

# Command to run the application
CMD ["./e-commerce"]
