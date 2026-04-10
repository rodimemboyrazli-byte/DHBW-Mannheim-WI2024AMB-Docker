FROM golang:1.25-alpine

WORKDIR /app

COPY main.go ./

# Das Projekt hat kein go.mod im Repo; wir initialisieren es im Build-Kontext.
RUN go mod init calculator && go mod tidy && go build -o server main.go

EXPOSE 5000

CMD ["./server"]
