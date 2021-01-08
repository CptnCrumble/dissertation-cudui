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

type hadsForm struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Tense            string
	Enjoy            string
	Fright           string
	Laugh            string
	Worry            string
	Cheer            string
	Ease             string
	Slow             string
	Butterfly        string
	Interest         string
	Restless         string
	Forward          string
	Panic            string
	Book             string
}

func newHads(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newhads.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, formData{false, patients})
				return
			}

			pid, _ := strconv.Atoi(r.FormValue("pid"))
			aNumber := r.FormValue("assessment_number")
			aDate := r.FormValue("assessment_date")
			tense := r.FormValue("tense")
			enjoy := r.FormValue("enjoy")
			fright := r.FormValue("fright")
			laugh := r.FormValue("laugh")
			worry := r.FormValue("worry")
			cheer := r.FormValue("cheer")
			ease := r.FormValue("ease")
			slow := r.FormValue("slow")
			butterfly := r.FormValue("butterfly")
			interest := r.FormValue("interest")
			restless := r.FormValue("restless")
			forward := r.FormValue("forward")
			panic := r.FormValue("panic")
			book := r.FormValue("book")

			j, _ := json.Marshal(hadsForm{pid, aNumber, aDate, tense, enjoy, fright, laugh, worry, cheer, ease, slow, butterfly, interest, restless, forward, panic, book})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_hads", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new hads with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newHads() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newHads() recieved response code of %d", response.StatusCode))
				b, _ := ioutil.ReadFile("./static/badsubmission.html")
				page := string(b)
				fmt.Fprintf(w, page)
			}
		}

	}
}
