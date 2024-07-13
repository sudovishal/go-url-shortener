# Build stage
FROM golang:1.20-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Install Node.js and npm for Tailwind CSS
RUN apk add --no-cache nodejs npm

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Install Tailwind CSS
RUN npm install -D tailwindcss

# Build Tailwind CSS
RUN npx tailwindcss -i ./input.css -o ./static/output.css

# Build the Go app
RUN go build -o main ./cmd/server

# Final stage
FROM alpine:latest

# Install dependencies for TursoDB client
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]