FROM golang:1.23.5-alpine as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./review ./review
COPY ./proto ./proto

WORKDIR /app/review
RUN go build -o review_service ./cmd/review/main.go

EXPOSE 50053
CMD ["./review_service"]
