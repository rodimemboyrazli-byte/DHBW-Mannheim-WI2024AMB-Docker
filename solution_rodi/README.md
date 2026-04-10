# solution_rodi

Drei einfache Calculator-APIs (Go, JavaScript, Python), die zusammen per Docker Compose laufen.

## Voraussetzungen

- Docker Desktop oder Docker Engine
- optional: `curl` fuer Tests

## Schnellstart

Im Ordner `solution_rodi`:

```bash
docker compose up -d --build
docker compose ps
```

Erreichbare Services:

- Go: `http://localhost:5001` (Host `5001` -> Container `5000`)
- JavaScript: `http://localhost:3000`
- Python: `http://localhost:4000`

Stoppen:

```bash
docker compose down
```

## API-Uebersicht (endpoints)

### Go (`localhost:5001`)

- `GET /sqrt/:x`
- `GET /mod/:a/:b`
- `GET /pow/:a/:b`
- `GET /log/:x`

### JavaScript (`localhost:3000`)

- `GET /add/:a/:b`
- `GET /substract/:a/:b`
- `GET /multiply/:a/:b`
- `GET /divide/:a/:b`

### Python (`localhost:4000`)

- `GET /sin/<x>`
- `GET /arcsin/<x>`
- `GET /cos/<x>`
- `GET /arccos/<x>`
- `GET /tan/<x>`
- `GET /arctan/<x>`

## Kurze Tests

```bash
curl -i "http://localhost:5001/sqrt/9"
curl -i "http://localhost:5001/log/-1"

curl -i "http://localhost:3000/add/2/3"
curl -i "http://localhost:3000/divide/10/0"

curl -i "http://localhost:4000/sin/1"
curl -i "http://localhost:4000/arccos/2"
```

Erwartete HTTP-Codes:

- Erfolg: `200`
- Fehlerfaelle: Go `400`, JavaScript `409`, Python `400`

## Validierung und Troubleshooting

Compose-Datei pruefen:

```bash
docker compose config
```

Bei Portkonflikt ("address already in use"):

```bash
docker compose down
docker ps
lsof -nP -iTCP:5001 -sTCP:LISTEN
```
