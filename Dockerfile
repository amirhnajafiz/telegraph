# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

# switch to root user
USER root

# work directory
WORKDIR /app

# copy go mod and go sum file
COPY go.mod go.sum ./

# copy all files
COPY . .

# go into cmd
WORKDIR /app/cmd

# build the file
RUN go build -o ./telegraph

# move to app directory
WORKDIR /app

# copy the build file into the root
RUN mv ./cmd/telegraph ./telegraph

# run the application
CMD ["./telegraph"]
