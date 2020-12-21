package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type patient struct {
	Pid       int
	Fname     string
	Sname     string
	Gender    string
	Diagnosis int
}

func newPatient(url string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		patients := getPids(url)

		tmpl := template.Must(template.ParseFiles("forms/newuser.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, formData{false, patients})
			return
		}

		pid, _ := strconv.Atoi(r.FormValue("pid"))
		gender := r.FormValue("gender")
		diagnosis, _ := strconv.Atoi(r.FormValue("diagnosis"))

		j, _ := json.Marshal(patient{pid, "anon", "anon", gender, diagnosis})

		fmt.Print(string(j))

		body := bytes.NewBuffer(j)
		api := fmt.Sprintf("%s/new_patient", url)
		response, err := http.Post(api, "application/json", body)

		if err != nil {
			fmt.Print("POST request failed")
			panic(err)
		}

		if response.StatusCode == 200 {
			tmpl.Execute(w, struct{ Success bool }{true})
		} else {
			fmt.Printf("non-200 response code")
		}
	}
}
