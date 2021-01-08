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

type pdq8Form struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Pdq1             string
	Pdq2             string
	Pdq3             string
	Pdq4             string
	Pdq5             string
	Pdq6             string
	Pdq7             string
	Pdq8             string
}

func newpdq8(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newpdq8.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, formData{false, patients})
				return
			}

			pid, _ := strconv.Atoi(r.FormValue("pid"))
			aNumber := r.FormValue("assessment_number")
			aDate := r.FormValue("assessment_date")
			p1 := r.FormValue("1")
			p2 := r.FormValue("2")
			p3 := r.FormValue("3")
			p4 := r.FormValue("4")
			p5 := r.FormValue("5")
			p6 := r.FormValue("6")
			p7 := r.FormValue("7")
			p8 := r.FormValue("8")

			j, _ := json.Marshal(pdq8Form{pid, aNumber, aDate, p1, p2, p3, p4, p5, p6, p7, p8})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_pdq8", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new pdq8 with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newPdq8() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newPdq8() recieved response code of %d", response.StatusCode))
				b, _ := ioutil.ReadFile("./static/badsubmission.html")
				page := string(b)
				fmt.Fprintf(w, page)
			}
		}

	}
}
