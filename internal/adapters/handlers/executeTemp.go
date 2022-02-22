package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func ExecTemp(PathHTML, NameHTML string, Error string, w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(PathHTML)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	err = tmpl.ExecuteTemplate(w, NameHTML, Error)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func Err(Str string, Status int, w http.ResponseWriter, r *http.Request) {
	// Info := Error{Str, Status}
	w.WriteHeader(Status)
	ExecTemp("templates/error.html", "error.html", "", w, r)
}
