Setup: 
1. `go mod download`
2. `go mod verify`
3. `go mod tidy`
4. `go build -o app`

``` Dockerfile
COPY --from=builder /app/app /app
```

Start: `/app`

Server läuft auf `localhost:5000`
