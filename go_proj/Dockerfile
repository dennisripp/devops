# Use an official Golang runtime as a parent image
FROM golang:1.22.2-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app as a static binary
RUN go mod tidy && \
    go build -o main .

# Use a small Alpine Linux image for the final stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./main"]

