# Base image
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN go build -o file-modification-tracker ./cmd/service

# Expose HTTP port
EXPOSE 8080

# Command to run the executable
CMD ["./file-modification-tracker"]
