package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

type userLogin struct {
	Username string
	Password string
}

func login(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// do the login phase
		tmpl := template.Must(template.ParseFiles("forms/login.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, formData{false, make([]int, 0)})
			return
		}

		uname := r.FormValue("uname")
		pword := r.FormValue("pword")

		debug := userLogin{Username: uname, Password: pword}

		j, _ := json.Marshal(debug)
		body := bytes.NewBuffer(j)
		api := fmt.Sprintf("%s/login", url)
		response, err := http.Post(api, "application/json", body)
		redisLogger(fmt.Sprintf("Login attempt for %s ", uname))

		if err != nil {
			redisLogger(fmt.Sprintf("Login attempt failed -- %s", err.Error()))
		}

		if response.StatusCode == 200 {
			// issue cookie & make .Success
			authkey := authKey()
			// expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "PCP", Value: authkey}
			http.SetCookie(w, &cookie)

			tmpl.Execute(w, struct{ Success bool }{true})
		} else {
			redisLogger(fmt.Sprintf("Login recieved response code of %d", response.StatusCode))
			// handle bad login attempt
		}

	}
}
