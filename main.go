package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	pgAdaptor := os.Getenv("URL_PG_ADAPTOR")

	r := mux.NewRouter()
	r.HandleFunc("/", landingPage())
	r.HandleFunc("/new_user", newPatient(pgAdaptor))
	r.HandleFunc("/new_nms", newNms(pgAdaptor))
	serve(r)
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
