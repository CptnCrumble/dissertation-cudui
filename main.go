package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	pgHost := os.Getenv("PG_ADAPTER_SERVICE_SERVICE_HOST")
	pgPort := os.Getenv("PG_ADAPTER_SERVICE_SERVICE_PORT")
	pgAdaptor := fmt.Sprintf("http://%s:%s", pgHost, pgPort)
	// needs a loop to wait for PG-adaptor to be up
	// ^ perhaps use k8's config to determine this

	r := mux.NewRouter()
	r.HandleFunc("/", landingPage(pgAdaptor))
	r.HandleFunc("/login", login(pgAdaptor))
	r.HandleFunc("/new_patient", newPatient(pgAdaptor))
	r.HandleFunc("/new_carer", newCarer(pgAdaptor))
	r.HandleFunc("/new_nms", newNms(pgAdaptor))
	r.HandleFunc("/new_updrs", newUpdrs(pgAdaptor))
	r.HandleFunc("/new_pdq39", newpdq39(pgAdaptor))
	r.HandleFunc("/new_pdqc", newpdqc(pgAdaptor))
	r.HandleFunc("/new_pdq8", newpdq8(pgAdaptor))
	r.HandleFunc("/new_hads", newHads(pgAdaptor))
	r.HandleFunc("/new_pdss", newPdss(pgAdaptor))
	r.HandleFunc("/new_pkg", newPkg(pgAdaptor))
	r.HandleFunc("/spreadsheets", spreadSheets(pgAdaptor))

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

func authCheck(r http.Request) bool {
	cookie, _ := r.Cookie("PCP")
	if cookie != nil {
		authkey := authKey()
		if cookie.Value == authkey {
			return true
		}
	}

	return false
}

func authKey() string {
	redisLocation := fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST"))
	conn, err := redis.Dial("tcp", redisLocation)
	if err != nil {
		log.Print(err)
		return "654651321deadredis6549876513213x"
	}

	// backup in case redis is down for pro-longed period
	// var value string

	value, err2 := redis.String(conn.Do("GET", "authkey"))
	if err2 != nil {
		log.Print(err2)
		return "654651321deadredis6549876513213x"
	}

	return value
}

func landingPage(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if authCheck(*r) {
			b, _ := ioutil.ReadFile("./static/index.html")
			page := string(b)
			fmt.Fprintf(w, page)
		} else {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		}
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

func getCids(url string) []int {
	c, err1 := http.Get(fmt.Sprintf("%s/carers", url))
	if err1 != nil {
		redisLogger(fmt.Sprintf("get all cids failed -- %s", err1.Error()))
		// return large array to prevent submission
		x := make([]int, 999)
		for i, v := range x {
			x[i] = i + v
		}
		return x
	}

	var cids []int
	json.NewDecoder(c.Body).Decode(&cids)
	return cids
}
