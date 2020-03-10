FROM golang:1.13 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy and download dependency using go mod
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o main .

FROM alpine:latest

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run when starting the container
ENTRYPOINT ["/main"]
CMD ["--help"]