package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type authResource struct{}

func (rs authResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", rs.Create)    // POST /v1/login - authorize user via login and password
	r.Post("/logout", rs.Create)   // POST /v1/logout - unauthorize user via token
	r.Post("/validate", rs.Create) // POST /v1/validate - validate tokens

	return r
}

func (rs authResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("."))
}
