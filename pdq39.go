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

type pdq39Form struct {
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
	Pdq9             string
	Pdq10            string
	Pdq11            string
	Pdq12            string
	Pdq13            string
	Pdq14            string
	Pdq15            string
	Pdq16            string
	Pdq17            string
	Pdq18            string
	Pdq19            string
	Pdq20            string
	Pdq21            string
	Pdq22            string
	Pdq23            string
	Pdq24            string
	Pdq25            string
	Pdq26            string
	Pdq27            string
	Pdq28            string
	Pdq29            string
	Pdq30            string
	Pdq31            string
	Pdq32            string
	Pdq33            string
	Pdq34            string
	Pdq35            string
	Pdq36            string
	Pdq37            string
	Pdq38            string
	Pdq39            string
}

func newpdq39(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authCheck(*r) {
			http.Redirect(w, r, "/login", http.StatusNetworkAuthenticationRequired)
		} else {
			patients := getPids(url)

			tmpl := template.Must(template.ParseFiles("forms/newpdq39.html"))
			if r.Method != http.MethodPost {
				tmpl.Execute(w, formData{false, patients})
				return
			}

			pid, _ := strconv.Atoi(r.FormValue("pid"))
			aNumber := r.FormValue("assessment_number")
			aDate := r.FormValue("assessment_date")
			// p1, _ := strconv.Atoi(r.FormValue("1"))
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
			p30 := r.FormValue("30")
			p31 := r.FormValue("31")
			p32 := r.FormValue("32")
			p33 := r.FormValue("33")
			p34 := r.FormValue("34")
			p35 := r.FormValue("35")
			p36 := r.FormValue("36")
			p37 := r.FormValue("37")
			p38 := r.FormValue("38")
			p39 := r.FormValue("39")

			j, _ := json.Marshal(pdq39Form{pid, aNumber, aDate, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18, p19, p20, p21, p22, p23, p24, p25, p26, p27, p28, p29, p30, p31, p32, p33, p34, p35, p36, p37, p38, p39})

			// POST to pgAdaptor
			body := bytes.NewBuffer(j)
			api := fmt.Sprintf("%s/new_pdq39", url)
			response, err := http.Post(api, "application/json", body)
			redisLogger(fmt.Sprintf("POSTed new pdq39 with %s", string(j)))

			if err != nil {
				redisLogger(fmt.Sprintf("newPdq39() POST failed -- %s", err.Error()))
			}

			if response.StatusCode == 200 {
				tmpl.Execute(w, struct{ Success bool }{true})
			} else {
				redisLogger(fmt.Sprintf("newPdq39() recieved response code of %d", response.StatusCode))
				b, _ := ioutil.ReadFile("./static/badsubmission.html")
				page := string(b)
				fmt.Fprintf(w, page)
			}
		}

	}
}
