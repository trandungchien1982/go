# Go QR Barcode API

GET `/barcode?code=Hello&imageSize=200` returns PNG (`Content-Type: image/png`).

## Defaults
- `WORKING_PORT=9092`
- `TYPE_CODE=QR_CODE` (only supported type)
- `IMAGE_SIZE=200`

`imageSize` overrides `IMAGE_SIZE` per request. Valid dimensions: 64–2048 pixels (square). `code` is required, max 2048 UTF-8 bytes. URL-encode query content, especially `&`, `+`, and `#`.

## Run locally (Go 1.22+)
```bash
go mod download
go run ./cmd/server
curl -f 'http://localhost:9092/barcode?code=Hello%20World&imageSize=300' -o qr.png
curl -f http://localhost:9092/health
```

## Docker Compose
```bash
cp .env.example .env
docker compose up --build -d
```
The default port is bound to localhost only. To expose publicly, edit `ports` and add appropriate network protections.

## Ubuntu native binary
```bash
chmod +x scripts/build-ubuntu.sh
./scripts/build-ubuntu.sh
./dist/barcode-server-linux-amd64
```
For ARM64: `GOARCH=arm64 ./scripts/build-ubuntu.sh`.

## HTTP responses
- 200: PNG
- 400: missing/invalid `code`, invalid `imageSize`, or QR payload too large
- 405: unsupported HTTP method

`TYPE_CODE` currently accepts `QR_CODE` only; other values cause startup failure.

Note: dependency download requires network access to proxy.golang.org on the first build.
