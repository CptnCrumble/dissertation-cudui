package main

import (
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
	Nms1             bool
	Nms2             bool
	Nms3             bool
	Nms4             bool
	Nms5             bool
	Nms6             bool
	Nms7             bool
	Nms8             bool
	Nms9             bool
	Nms10            bool
	Nms11            bool
	Nms12            bool
	Nms13            bool
	Nms14            bool
	Nms15            bool
	Nms16            bool
	Nms17            bool
	Nms18            bool
	Nms19            bool
	Nms20            bool
	Nms21            bool
	Nms22            bool
	Nms23            bool
	Nms24            bool
	Nms25            bool
	Nms26            bool
	Nms27            bool
	Nms28            bool
	Nms29            bool
	Nms30            bool
	Nms31            bool
}

func newNms(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("forms/newnms.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		pid, _ := strconv.Atoi(r.FormValue("pid"))
		aNumber := r.FormValue("assessment_number")
		aDate := r.FormValue("assessment_date")
		nms1, _ := strconv.ParseBool(r.FormValue("nms1"))
		nms2, _ := strconv.ParseBool(r.FormValue("nms2"))
		nms3, _ := strconv.ParseBool(r.FormValue("nms3"))
		nms4, _ := strconv.ParseBool(r.FormValue("nms4"))
		nms5, _ := strconv.ParseBool(r.FormValue("nms5"))
		nms6, _ := strconv.ParseBool(r.FormValue("nms6"))
		nms7, _ := strconv.ParseBool(r.FormValue("nms7"))
		nms8, _ := strconv.ParseBool(r.FormValue("nms8"))
		nms9, _ := strconv.ParseBool(r.FormValue("nms9"))
		nms10, _ := strconv.ParseBool(r.FormValue("nms10"))
		nms11, _ := strconv.ParseBool(r.FormValue("nms11"))
		nms12, _ := strconv.ParseBool(r.FormValue("nms12"))
		nms13, _ := strconv.ParseBool(r.FormValue("nms13"))
		nms14, _ := strconv.ParseBool(r.FormValue("nms14"))
		nms15, _ := strconv.ParseBool(r.FormValue("nms15"))
		nms16, _ := strconv.ParseBool(r.FormValue("nms16"))
		nms17, _ := strconv.ParseBool(r.FormValue("nms17"))
		nms18, _ := strconv.ParseBool(r.FormValue("nms18"))
		nms19, _ := strconv.ParseBool(r.FormValue("nms19"))
		nms20, _ := strconv.ParseBool(r.FormValue("nms20"))
		nms21, _ := strconv.ParseBool(r.FormValue("nms21"))
		nms22, _ := strconv.ParseBool(r.FormValue("nms22"))
		nms23, _ := strconv.ParseBool(r.FormValue("nms23"))
		nms24, _ := strconv.ParseBool(r.FormValue("nms24"))
		nms25, _ := strconv.ParseBool(r.FormValue("nms25"))
		nms26, _ := strconv.ParseBool(r.FormValue("nms26"))
		nms27, _ := strconv.ParseBool(r.FormValue("nms27"))
		nms28, _ := strconv.ParseBool(r.FormValue("nms28"))
		nms29, _ := strconv.ParseBool(r.FormValue("nms29"))
		nms30, _ := strconv.ParseBool(r.FormValue("nms30"))
		nms31, _ := strconv.ParseBool(r.FormValue("nms31"))

		j, _ := json.Marshal(nmsForm{pid, aNumber, aDate, nms1, nms2, nms3, nms4, nms5, nms6, nms7, nms8, nms9, nms10, nms11, nms12, nms13, nms14, nms15, nms16, nms17, nms18, nms19, nms20, nms21, nms22, nms23, nms24, nms25, nms26, nms27, nms28, nms29, nms30, nms31})
		fmt.Print(string(j))

		// POST to pgAdaptor

	}
}
