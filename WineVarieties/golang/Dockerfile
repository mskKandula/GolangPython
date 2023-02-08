# Stage 1 as builder
FROM golang:1.18-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app/winevarieties/golang

# Copy Go Application Dependencies
COPY ./golang/go.mod .
COPY ./golang/go.sum .

# To keep only actual dependencies
RUN go mod tidy

# Download dependencies 
RUN go mod download

# Copy code files
COPY ./golang/ .

# Build the Go app
RUN go build -o client .

########### Stage 2 ###########

#Define the running image
FROM alpine:3.17

RUN apk update && \
    apk upgrade

# Set the Current Working Directory inside the container
WORKDIR /app/winevarieties/golang

# Copy built binary application from 'builder' image
COPY --from=builder /app/winevarieties/golang/client .

# This container exposes port 8080 to the outside world
EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["./client"]
