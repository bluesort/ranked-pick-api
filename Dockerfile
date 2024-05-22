FROM golang:1.22-alpine
RUN apk add build-base
WORKDIR /go/app/ranked-pick-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/bin/ranked-pick-api github.com/carterjackson/ranked-pick-api/cmd/api
