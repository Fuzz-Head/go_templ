package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Fuzz-Head/go_htmx_tailwind/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

  if err := godotenv.Load(); err != nil {
    log.Fatal(err)
  }

  router := chi.NewMux()
  
  router.Handle("/*", public())

  router.Get("/foo", handlers.Make(handlers.HandleFoo))

  listenAddr := os.Getenv("LISTEN_ADDR")
  slog.Info("HTTP server started", "listenAddr", listenAddr)
  http.ListenAndServe(listenAddr, router)
}

