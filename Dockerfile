FROM golang:1.22-alpine AS build
RUN apk add build-base
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go env -w GOCACHE=/go-cache
COPY . .
RUN --mount=type=cache,target=/go-cache go build -o /usr/bin/ranked-pick-api github.com/carterjackson/ranked-pick-api/cmd/api
