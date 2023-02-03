# Use the latest alpine-based Go image as the builder stage
FROM golang:alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go source files
COPY . .

# We need to add some test to run before the build
#RUN go test ./...

# Build the Go binary
RUN go build -o sugarcity-bot ./cmd/chat-bot

# Use the Alpine Linux image as the runtime stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/chat-bot .

# Set the command to run the binary
CMD ["./chat-bot"]
