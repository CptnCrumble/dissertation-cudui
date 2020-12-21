package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	pgAdaptor := os.Getenv("URL_PG_ADAPTOR")
	// needs a loop to wait for PG-adaptor to be up
	p, err1 := http.Get(fmt.Sprintf("%s/users", pgAdaptor))
	if err1 != nil {
		fmt.Print("get all pids failed")
	}

	var pids []int
	json.NewDecoder(p.Body).Decode(&pids)
	fmt.Print(pids)

	r := mux.NewRouter()
	r.HandleFunc("/", landingPage())
	r.HandleFunc("/new_patient", newPatient(pgAdaptor))
	r.HandleFunc("/new_nms", newNms(pgAdaptor, pids))
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
