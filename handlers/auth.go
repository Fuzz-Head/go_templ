package handlers

import (
  "net/http"

  "github.com/Fuzz-Head/go_htmx_tailwind/views/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
  return Render(w, r, auth.Login())
}
