# syntax=docker/dockerfile:1
FROM golang:1.21-alpine as builder
ENV GO111MODULE=on
LABEL authors="danflanagan"

ARG PORT=8080

# Set destination for COPY
WORKDIR /app

# Set destination for COPY
COPY go.mod go.sum ./

RUN go mod tidy

# Copy the source code
COPY explorer-api/ .
COPY lib/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

WORKDIR /app

CMD ["./app"]
EXPOSE $PORT