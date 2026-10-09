FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/barcode-server ./cmd/server

FROM scratch

COPY --from=build /out/barcode-server /barcode-server

EXPOSE 9092

ENTRYPOINT ["/barcode-server"]
