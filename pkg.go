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

type pkgForm struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	DurationDays     int
	Bks              float64
	Dks              float64
	Fds              float64
	Pti              float64
	Ptt              float64
}

func newPkg(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newpkg.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, formData{false, patients})
				return
			}

			pid, _ := strconv.Atoi(r.FormValue("pid"))
			aNumber := r.FormValue("assessment_number")
			aDate := r.FormValue("assessment_date")
			duration, _ := strconv.Atoi(r.FormValue("duration"))
			bks, _ := strconv.ParseFloat(r.FormValue("bks"), 64)
			dks, _ := strconv.ParseFloat(r.FormValue("dks"), 64)
			fds, _ := strconv.ParseFloat(r.FormValue("fds"), 64)
			pti, _ := strconv.ParseFloat(r.FormValue("pti"), 64)
			ptt, _ := strconv.ParseFloat(r.FormValue("ptt"), 64)

			j, _ := json.Marshal(pkgForm{pid, aNumber, aDate, duration, bks, dks, fds, pti, ptt})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_pkg", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new pkg with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newPkg() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newPkg() recieved response code of %d", response.StatusCode))
				b, _ := ioutil.ReadFile("./static/badsubmission.html")
				page := string(b)
				fmt.Fprintf(w, page)
			}
		}

	}
}
