package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type pdq39Form struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Pdq1             int
	Pdq2             int
	Pdq3             int
	Pdq4             int
	Pdq5             int
	Pdq6             int
	Pdq7             int
	Pdq8             int
	Pdq9             int
	Pdq10            int
	Pdq11            int
	Pdq12            int
	Pdq13            int
	Pdq14            int
	Pdq15            int
	Pdq16            int
	Pdq17            int
	Pdq18            int
	Pdq19            int
	Pdq20            int
	Pdq21            int
	Pdq22            int
	Pdq23            int
	Pdq24            int
	Pdq25            int
	Pdq26            int
	Pdq27            int
	Pdq28            int
	Pdq29            int
	Pdq30            int
	Pdq31            int
	Pdq32            int
	Pdq33            int
	Pdq34            int
	Pdq35            int
	Pdq36            int
	Pdq37            int
	Pdq38            int
	Pdq39            int
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
			p1, _ := strconv.Atoi(r.FormValue("1"))
			p2, _ := strconv.Atoi(r.FormValue("2"))
			p3, _ := strconv.Atoi(r.FormValue("3"))
			p4, _ := strconv.Atoi(r.FormValue("4"))
			p5, _ := strconv.Atoi(r.FormValue("5"))
			p6, _ := strconv.Atoi(r.FormValue("6"))
			p7, _ := strconv.Atoi(r.FormValue("7"))
			p8, _ := strconv.Atoi(r.FormValue("8"))
			p9, _ := strconv.Atoi(r.FormValue("9"))
			p10, _ := strconv.Atoi(r.FormValue("10"))
			p11, _ := strconv.Atoi(r.FormValue("11"))
			p12, _ := strconv.Atoi(r.FormValue("12"))
			p13, _ := strconv.Atoi(r.FormValue("13"))
			p14, _ := strconv.Atoi(r.FormValue("14"))
			p15, _ := strconv.Atoi(r.FormValue("15"))
			p16, _ := strconv.Atoi(r.FormValue("16"))
			p17, _ := strconv.Atoi(r.FormValue("17"))
			p18, _ := strconv.Atoi(r.FormValue("18"))
			p19, _ := strconv.Atoi(r.FormValue("19"))
			p20, _ := strconv.Atoi(r.FormValue("20"))
			p21, _ := strconv.Atoi(r.FormValue("21"))
			p22, _ := strconv.Atoi(r.FormValue("22"))
			p23, _ := strconv.Atoi(r.FormValue("23"))
			p24, _ := strconv.Atoi(r.FormValue("24"))
			p25, _ := strconv.Atoi(r.FormValue("25"))
			p26, _ := strconv.Atoi(r.FormValue("26"))
			p27, _ := strconv.Atoi(r.FormValue("27"))
			p28, _ := strconv.Atoi(r.FormValue("28"))
			p29, _ := strconv.Atoi(r.FormValue("29"))
			p30, _ := strconv.Atoi(r.FormValue("30"))
			p31, _ := strconv.Atoi(r.FormValue("31"))
			p32, _ := strconv.Atoi(r.FormValue("32"))
			p33, _ := strconv.Atoi(r.FormValue("33"))
			p34, _ := strconv.Atoi(r.FormValue("34"))
			p35, _ := strconv.Atoi(r.FormValue("35"))
			p36, _ := strconv.Atoi(r.FormValue("36"))
			p37, _ := strconv.Atoi(r.FormValue("37"))
			p38, _ := strconv.Atoi(r.FormValue("38"))
			p39, _ := strconv.Atoi(r.FormValue("39"))

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
			}
		}

	}
}
