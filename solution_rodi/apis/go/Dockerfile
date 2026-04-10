FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/app /app/app
EXPOSE 5000
CMD ["/app/app"]

