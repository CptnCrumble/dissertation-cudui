package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type pdssForm struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Pdss1            string
	Pdss2            string
	Pdss3            string
	Pdss4            string
	Pdss5            string
	Pdss6            string
	Pdss7            string
	Pdss8            string
	Pdss9            string
	Pdss10           string
	Pdss11           string
	Pdss12           string
	Pdss13           string
	Pdss14           string
	Pdss15           string
}

func newPdss(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newpdss.html"))
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
			p9 := r.FormValue("9")
			p10 := r.FormValue("10")
			p11 := r.FormValue("11")
			p12 := r.FormValue("12")
			p13 := r.FormValue("13")
			p14 := r.FormValue("14")
			p15 := r.FormValue("15")

			j, _ := json.Marshal(pdssForm{pid, aNumber, aDate, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_pdss", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new pdss with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newPdss() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newPdss() recieved response code of %d", response.StatusCode))
			}
		}

	}
}
