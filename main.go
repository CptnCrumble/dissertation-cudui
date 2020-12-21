package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	pgAdaptor := os.Getenv("URL_PG_ADAPTOR")
	// needs a loop to wait for PG-adaptor to be up
	// ^ perhaps use k8's config to determine this

	r := mux.NewRouter()
	r.HandleFunc("/", landingPage())
	r.HandleFunc("/new_patient", newPatient(pgAdaptor))
	r.HandleFunc("/new_nms", newNms(pgAdaptor))
	serve(r)
}

type formData struct {
	Success bool
	Pids    []int
}

func serve(router *mux.Router) {
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}

func redisLogger(message string) {
	redisLocation := fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST"))
	conn, err := redis.Dial("tcp", redisLocation)

	if err != nil {
		log.Print(err)
		return
	}
	defer conn.Close()

	t := time.Now()
	log2go := fmt.Sprintf("CUDUI -- %s -- %s", t.Format("2006-01-02 15:04:05"), message)
	_, err = conn.Do("LPUSH", "logs", log2go)

	if err != nil {
		log.Print("Couldn't log to redis")
		log.Print(log2go)
	} else {
		log.Print("logged to redis")
	}
}

func landingPage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "landing page")
	}
}

func getPids(url string) []int {
	p, err1 := http.Get(fmt.Sprintf("%s/users", url))
	if err1 != nil {
		redisLogger(fmt.Sprintf("get all pids failed -- %s", err1.Error()))
		// return large array to prevent submission
		x := make([]int, 999)
		for i, v := range x {
			x[i] = i + v
		}
		return x
	}

	var pids []int
	json.NewDecoder(p.Body).Decode(&pids)
	return pids
}
