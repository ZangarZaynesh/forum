package handlers

import "net/http"

type User interface {
	Register(router *http.ServeMux)
}
