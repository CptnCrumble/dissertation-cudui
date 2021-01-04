package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type nmsForm struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Nms1             string
	Nms2             string
	Nms3             string
	Nms4             string
	Nms5             string
	Nms6             string
	Nms7             string
	Nms8             string
	Nms9             string
	Nms10            string
	Nms11            string
	Nms12            string
	Nms13            string
	Nms14            string
	Nms15            string
	Nms16            string
	Nms17            string
	Nms18            string
	Nms19            string
	Nms20            string
	Nms21            string
	Nms22            string
	Nms23            string
	Nms24            string
	Nms25            string
	Nms26            string
	Nms27            string
	Nms28            string
	Nms29            string
	Nms30            string
	Nms31            string
}

func newNms(url string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newnms.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, formData{false, patients})
				return
			}

			pid, _ := strconv.Atoi(r.FormValue("pid"))
			aNumber := r.FormValue("assessment_number")
			aDate := r.FormValue("assessment_date")
			nms1 := r.FormValue("nms1")
			nms2 := r.FormValue("nms2")
			nms3 := r.FormValue("nms3")
			nms4 := r.FormValue("nms4")
			nms5 := r.FormValue("nms5")
			nms6 := r.FormValue("nms6")
			nms7 := r.FormValue("nms7")
			nms8 := r.FormValue("nms8")
			nms9 := r.FormValue("nms9")
			nms10 := r.FormValue("nms10")
			nms11 := r.FormValue("nms11")
			nms12 := r.FormValue("nms12")
			nms13 := r.FormValue("nms13")
			nms14 := r.FormValue("nms14")
			nms15 := r.FormValue("nms15")
			nms16 := r.FormValue("nms16")
			nms17 := r.FormValue("nms17")
			nms18 := r.FormValue("nms18")
			nms19 := r.FormValue("nms19")
			nms20 := r.FormValue("nms20")
			nms21 := r.FormValue("nms21")
			nms22 := r.FormValue("nms22")
			nms23 := r.FormValue("nms23")
			nms24 := r.FormValue("nms24")
			nms25 := r.FormValue("nms25")
			nms26 := r.FormValue("nms26")
			nms27 := r.FormValue("nms27")
			nms28 := r.FormValue("nms28")
			nms29 := r.FormValue("nms29")
			nms30 := r.FormValue("nms30")
			nms31 := r.FormValue("nms31")

			j, _ := json.Marshal(nmsForm{pid, aNumber, aDate, nms1, nms2, nms3, nms4, nms5, nms6, nms7, nms8, nms9, nms10, nms11, nms12, nms13, nms14, nms15, nms16, nms17, nms18, nms19, nms20, nms21, nms22, nms23, nms24, nms25, nms26, nms27, nms28, nms29, nms30, nms31})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_nms", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new nms with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newNms() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newNms() recieved response code of %d", response.StatusCode))
			}
		}

	}
}
