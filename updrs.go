package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type updrsForm struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Speech           int
	Saliva           int
	Chewing          int
	Eating           int
	Dressing         int
	Hygiene          int
	Handwriting      int
	Hobbies          int
	Turning          int
}

func newUpdrs(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newupdrs.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, formData{false, patients})
				return
			}

			pid, _ := strconv.Atoi(r.FormValue("pid"))
			aNumber := r.FormValue("assessment_number")
			aDate := r.FormValue("assessment_date")
			speech, _ := strconv.Atoi(r.FormValue("speech"))
			saliva, _ := strconv.Atoi(r.FormValue("saliva"))
			chewing, _ := strconv.Atoi(r.FormValue("chewing"))
			eating, _ := strconv.Atoi(r.FormValue("eating"))
			dressing, _ := strconv.Atoi(r.FormValue("dressing"))
			hygiene, _ := strconv.Atoi(r.FormValue("hygiene"))
			handwriting, _ := strconv.Atoi(r.FormValue("handwriting"))
			hobbies, _ := strconv.Atoi(r.FormValue("hobbies"))
			turning, _ := strconv.Atoi(r.FormValue("turning"))

			j, _ := json.Marshal(updrsForm{pid, aNumber, aDate, speech, saliva, chewing, eating, dressing, hygiene, handwriting, hobbies, turning})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_updrs", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new updrs with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newUpdrs() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newUpdrs() recieved response code of %d", response.StatusCode))
			}
		}

	}
}
