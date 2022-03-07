package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func ExecTemp(Date interface{}, NameHTML string, w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/" + NameHTML)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	err = tmpl.ExecuteTemplate(w, NameHTML, Date)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func Err(Str string, Status int, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(Status)
	ExecTemp(Str, "error.html", w, r)
}
