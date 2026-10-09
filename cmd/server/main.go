package main

import (
 "errors"
 "fmt"
 "log"
 "net/http"
 "os"
 "strconv"
 "strings"
 "time"

 qrcode "github.com/skip2/go-qrcode"
)

func env(key, fallback string) string {
 if value := strings.TrimSpace(os.Getenv(key)); value != "" { return value }
 return fallback
}

func main() {
 port := env("WORKING_PORT", "9092")
 if n, err := strconv.Atoi(port); err != nil || n < 1 || n > 65535 { log.Fatal("WORKING_PORT must be between 1 and 65535") }
 codeType := strings.ToUpper(env("TYPE_CODE", "QR_CODE"))
 if codeType != "QR_CODE" { log.Fatalf("unsupported TYPE_CODE %q: only QR_CODE is supported", codeType) }
 defaultSize, err := strconv.Atoi(env("IMAGE_SIZE", "200"))
 if err != nil || defaultSize < 64 || defaultSize > 2048 { log.Fatal("IMAGE_SIZE must be between 64 and 2048") }
 mux := http.NewServeMux()
 mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "text/plain; charset=utf-8"); fmt.Fprint(w, "OK") })
 mux.HandleFunc("GET /barcode", func(w http.ResponseWriter, r *http.Request) {
  code := r.URL.Query().Get("code")
  if strings.TrimSpace(code) == "" { http.Error(w, "missing or empty query parameter: code", http.StatusBadRequest); return }
  if len(code) > 2048 { http.Error(w, "code exceeds 2048 bytes", http.StatusBadRequest); return }
  size := defaultSize
  if raw := r.URL.Query().Get("imageSize"); raw != "" {
   parsed, err := strconv.Atoi(raw)
   if err != nil || parsed < 64 || parsed > 2048 { http.Error(w, "imageSize must be an integer between 64 and 2048", http.StatusBadRequest); return }
   size = parsed
  }
  png, err := qrcode.Encode(code, qrcode.Medium, size)
  if err != nil { http.Error(w, "unable to encode QR code", http.StatusBadRequest); return }
  w.Header().Set("Content-Type", "image/png")
  w.Header().Set("Cache-Control", "no-store")
  w.Header().Set("Content-Length", strconv.Itoa(len(png)))
  _, _ = w.Write(png)
 })
 server := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5*time.Second, IdleTimeout: 60*time.Second}
 log.Printf("QR barcode server listening on :%s", port)
 if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) { log.Fatal(err) }
}
