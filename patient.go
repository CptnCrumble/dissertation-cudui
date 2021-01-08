package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type patient struct {
	Pid       int
	Fname     string
	Sname     string
	Gender    string
	Diagnosis int
}

func newPatient(url string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newuser.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, formData{false, patients})
				return
			}

			pid, _ := strconv.Atoi(r.FormValue("pid"))
			gender := r.FormValue("gender")
			diagnosis, _ := strconv.Atoi(r.FormValue("diagnosis"))

			j, _ := json.Marshal(patient{pid, "anon", "anon", gender, diagnosis})

			redisLogger(fmt.Sprintf("POSTed new patient with %s", string(j)))

			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_patient", url)
			response, err := http.Post(api, "application/json", body)

			if err != nil {
				redisLogger(fmt.Sprintf("newPatient POST request failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newPatient() recieved response code of %d", response.StatusCode))
				b, _ := ioutil.ReadFile("./static/badsubmission.html")
				page := string(b)
				fmt.Fprintf(w, page)
			}
		}

	}
}
