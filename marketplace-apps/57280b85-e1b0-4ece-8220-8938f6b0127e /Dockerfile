# Start from the official Go image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o app .

# Command to run the application
CMD ["./app"]
