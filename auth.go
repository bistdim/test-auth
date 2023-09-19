package main

import (
	"net/http"
	"os"

	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gopkg.in/yaml.v3"
)

type authResource struct{}
type ErrResponse struct {
	Err            error  `json:"-"`      // low-level runtime error
	HTTPStatusCode int    `json:"-"`      // http response status code
	StatusText     string `json:"status"` // user-level status message
}

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var ErrAnauthorized = &ErrResponse{HTTPStatusCode: 401, StatusText: "Invalid Login and Password"}

func (rs authResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", rs.Login)    // POST /v1/login - authorize user via login and password
	r.Post("/logout", rs.Login)   // POST /v1/logout - unauthorize user via token
	r.Post("/validate", rs.Login) // POST /v1/validate - validate tokens

	return r
}

func (rs authResource) Login(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("credentials.yml")
	if err != nil {
		processError(err)
	}

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		processError(err)
	}

	//w.Write([]byte(r.user))
	var u User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		processError(err)
	}

	if cfg.User.Login == u.User && cfg.User.Password == u.Password {
		return
	} else {
		render.Render(w, r, ErrAnauthorized)
		return
	}
}
