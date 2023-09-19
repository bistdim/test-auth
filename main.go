package main

import (
	"net/http"

	"os"

	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	User struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	} `yaml:"user"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/v1", authResource{}.Routes())

	http.ListenAndServe(":8080", r)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
