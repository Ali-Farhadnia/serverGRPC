# Start from golang base image
FROM golang:alpine as grpc_server

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Ali Farhadnia <ali.farhadnia.80@gmail.com>"

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .
ARG grpc_address
ARG grpc_network
ARG db_user
ARG db_password
ARG db_dbname
ARG db_sslmode
ARG db_port
ARG db_host
ENV grpc_address =$grpc_address grpc_network=$grpc_network db_user=$db_user db_password=$db_password db_dbname=$db_dbname db_sslmode=$db_sslmode db_port=$db_port db_host=$db_host

# Build the Go app
RUN go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=grpc_server /app/main .
COPY --from=grpc_server /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]