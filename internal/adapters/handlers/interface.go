package handlers

import "net/http"

type User interface {
	Register(router *http.ServeMux)
}

type Post interface {
	Register(router *http.ServeMux)
}
