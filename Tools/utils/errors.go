package utils

import (
	"net/http"
	"text/template"
)

func Error500(tm *template.Template, rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	tm.ExecuteTemplate(rw, "error.html", "Oops! Internal server Error")
}

func Error400(tm *template.Template, rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	tm.ExecuteTemplate(rw, "error.html", "Bad Request Code 400")
}

func Error404(tm *template.Template, rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotFound)
	tm.ExecuteTemplate(rw, "error.html", "Not Found Code 404")
}

func ErrorUrl(r *http.Request, s string) bool {
	return r.URL.Path != s
}
