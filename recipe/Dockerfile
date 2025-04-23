FROM golang:1.23.5-alpine as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./recipe ./recipe
COPY ./proto ./proto

WORKDIR /app/recipe
RUN go build -o recipe_service ./cmd/recipe/main.go

EXPOSE 8082
CMD ["./recipe_service"]