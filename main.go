package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", landingPage())
	// new user
	tmpl := template.Must(template.ParseFiles("newuser.html"))
	r.HandleFunc("/new_user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		role := r.FormValue("role")
		gender := r.FormValue("gender")
		diagnosis := r.FormValue("diagnosis")

		// replace with JSON post tp PG-adapter

		tmpl.Execute(w, struct{ Success bool }{true})
		fmt.Printf("%s %s %s", role, gender, diagnosis)

	})
	serve(r)
}

type client struct {
	Id        int
	Role      string
	Gender    string
	Diagnosis int
}

func serve(router *mux.Router) {
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}

func landingPage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "landing page")
	}
}
