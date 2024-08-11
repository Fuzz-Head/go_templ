package handlers

import (
	"net/http"

	"github.com/Fuzz-Head/go_htmx_tailwind/views/home"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
  // generally done in this way 
  // return foo.Index().Render(r.Context(), w)
  return Render(w, r, home.Index())

}
