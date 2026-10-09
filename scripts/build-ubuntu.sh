#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p dist
GOOS=linux GOARCH="${GOARCH:-amd64}" CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o "dist/barcode-server-linux-${GOARCH:-amd64}" ./cmd/server
printf 'Built: dist/barcode-server-linux-%s\n' "${GOARCH:-amd64}"
