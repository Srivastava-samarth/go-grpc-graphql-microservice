# Build stage
FROM golang:1.22-alpine as build

# Install build dependencies
RUN apk --no-cache add gcc g++ make ca-certificates

# Set the working directory
WORKDIR /go-grpc-graphql-microservice

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project (including account and catalog directories)
COPY . .

# Build the application
RUN GO111MODULE=on go build -o /go/bin/app ./catalog/cmd/catalog

# Final stage
FROM alpine:3.18

# Set the working directory
WORKDIR /usr/bin

# Copy the binary from the build stage
COPY --from=build /go/bin/app .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./app"]
