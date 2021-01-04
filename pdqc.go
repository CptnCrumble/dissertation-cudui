package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type pdqCForm struct {
	Cid              int
	AssessmentNumber string
	AssessmentDate   string
	Pdqc1            string
	Pdqc2            string
	Pdqc3            string
	Pdqc4            string
	Pdqc5            string
	Pdqc6            string
	Pdqc7            string
	Pdqc8            string
	Pdqc9            string
	Pdqc10           string
	Pdqc11           string
	Pdqc12           string
	Pdqc13           string
	Pdqc14           string
	Pdqc15           string
	Pdqc16           string
	Pdqc17           string
	Pdqc18           string
	Pdqc19           string
	Pdqc20           string
	Pdqc21           string
	Pdqc22           string
	Pdqc23           string
	Pdqc24           string
	Pdqc25           string
	Pdqc26           string
	Pdqc27           string
	Pdqc28           string
	Pdqc29           string
}

func newpdqc(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			carers := getCids(url)

			tmpl := template.Must(template.ParseFiles("forms/newpdqc.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, carerFormData{false, make([]int, 0), carers})
				return
			}

			cid, _ := strconv.Atoi(r.FormValue("cid"))
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
			p16 := r.FormValue("16")
			p17 := r.FormValue("17")
			p18 := r.FormValue("18")
			p19 := r.FormValue("19")
			p20 := r.FormValue("20")
			p21 := r.FormValue("21")
			p22 := r.FormValue("22")
			p23 := r.FormValue("23")
			p24 := r.FormValue("24")
			p25 := r.FormValue("25")
			p26 := r.FormValue("26")
			p27 := r.FormValue("27")
			p28 := r.FormValue("28")
			p29 := r.FormValue("29")

			j, _ := json.Marshal(pdqCForm{cid, aNumber, aDate, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18, p19, p20, p21, p22, p23, p24, p25, p26, p27, p28, p29})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_pdqc", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new pdqc with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newPdqc() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newPdqc() recieved response code of %d", response.StatusCode))
			}
		}

	}
}
