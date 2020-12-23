package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type carer struct {
	Cid          int
	Fname        string
	Sname        string
	Pid          int
	Relationship string
}

type carerFormData struct {
	Success bool
	Pids    []int
	Cids    []int
}

func newCarer(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		patients := getPids(url)
		carers := getCids(url)

		tmpl := template.Must(template.ParseFiles("forms/newcarer.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, carerFormData{false, patients, carers})
			return
		}

		pid, _ := strconv.Atoi(r.FormValue("pid"))
		cid, _ := strconv.Atoi(r.FormValue("cid"))
		rel := r.FormValue("relationship")

		j, _ := json.Marshal(carer{cid, "anon", "anon", pid, rel})
		body := bytes.NewBuffer(j)
		api := fmt.Sprintf("%s/new_carer", url)
		redisLogger(fmt.Sprintf("POSTed new carer with %s", string(j)))
		response, err := http.Post(api, "application/json", body)

		if err != nil {
			redisLogger(fmt.Sprintf("newCarer POST request failed -- %s", err.Error()))
		}

		if response.StatusCode == 200 {
			tmpl.Execute(w, struct{ Success bool }{true})
		} else {
			redisLogger(fmt.Sprintf("newCarer() recieved response code of %d", response.StatusCode))
		}
	}
}
