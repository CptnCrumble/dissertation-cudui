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
	PdqC1            int
	PdqC2            int
	PdqC3            int
	PdqC4            int
	PdqC5            int
	PdqC6            int
	PdqC7            int
	PdqC8            int
	PdqC9            int
	PdqC10           int
	PdqC11           int
	PdqC12           int
	PdqC13           int
	PdqC14           int
	PdqC15           int
	PdqC16           int
	PdqC17           int
	PdqC18           int
	PdqC19           int
	PdqC20           int
	PdqC21           int
	PdqC22           int
	PdqC23           int
	PdqC24           int
	PdqC25           int
	PdqC26           int
	PdqC27           int
	PdqC28           int
	PdqC29           int
}

func newpdqc(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		carers := getCids(url)

		tmpl := template.Must(template.ParseFiles("forms/newpdqc.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, carerFormData{false, make([]int, 0), carers})
			return
		}

		cid, _ := strconv.Atoi(r.FormValue("cid"))
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
