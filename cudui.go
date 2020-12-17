package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	pgAdaptor := os.Getenv("URL_PG_ADAPTOR")

	r := mux.NewRouter()
	r.HandleFunc("/", landingPage())
	r.HandleFunc("/new_user", newUser(pgAdaptor))
	// r.HandleFunc("/new_nms", newNms(pgAdaptor))
	serve(r)
}

type patient struct {
	Pid       int
	Fname     string
	Sname     string
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

func newUser(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("forms/newuser.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		pid, _ := strconv.Atoi(r.FormValue("pid"))
		gender := r.FormValue("gender")
		diagnosis, _ := strconv.Atoi(r.FormValue("diagnosis"))

		j, _ := json.Marshal(patient{pid, "anon", "anon", gender, diagnosis})

		fmt.Print(string(j))

		body := bytes.NewBuffer(j)
		api := fmt.Sprintf("%s/new_user", url)
		response, err := http.Post(api, "application/json", body)

		if err != nil {
			fmt.Print("POST request failed")
			panic(err)
		}

		if response.StatusCode == 200 {
			tmpl.Execute(w, struct{ Success bool }{true})
		} else {
			fmt.Printf("non-200 response code")
		}
	}
}

// func newNms(url string) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		tmpl := template.Must(template.ParseFiles("newuser.html"))
// 		if r.Method != http.MethodPost {
// 			tmpl.Execute(w, nil)
// 			return
// 		}
// 	}
// }
