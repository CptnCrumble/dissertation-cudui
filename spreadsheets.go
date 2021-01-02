package main

import (
	"html/template"
	"net/http"
	"os"
)

type spreadSheetPage struct {
	Pids           []int
	XlsxFactoryURL string
}

func spreadSheets(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)
			urlXf := os.Getenv("URL_XLSX_FACTORY")

			tmpl := template.Must(template.ParseFiles("static/spreadsheets.html"))
			tmpl.Execute(w, spreadSheetPage{patients, urlXf})
		}
	}
}
