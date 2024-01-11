# syntax=docker/dockerfile:1
FROM golang:1.21-alpine as builder
ENV GO111MODULE=on
LABEL authors="danflanagan"

ARG COMMIT_HASH
ENV COMMIT_HASH=${COMMIT_HASH}
ARG BUILD_DATE
ENV BUILD_DATE=${BUILD_DATE}
ARG ENV
ENV ENV=${ENV}
ARG VERSION
ENV VERSION=${VERSION}

ARG PORT=8080

# Set destination for COPY
WORKDIR /app

# Set destination for COPY
COPY go.mod go.sum ./

RUN go mod tidy

# Copy the source code
COPY . .

RUN echo ${COMMIT_HASH}
RUN echo ${BUILD_DATE}
RUN echo ${ENV}
RUN echo ${VERSION}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

WORKDIR /app

CMD ["./app"]
EXPOSE $PORT