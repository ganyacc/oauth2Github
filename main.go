package main

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubOauthConfig = &oauth2.Config{
	ClientID:     "27076bd1715186dfbda8",
	ClientSecret: "40ee42c67b067331468443e0b01f7ad8bf786112",
	Endpoint:     github.Endpoint,
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", startGithubOauth)
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, `<html>
	
	<body>
	<form action="/upload" method="post">

	<input type= "submit" value= "login with github">
	</body>
	
	</html>`)
}

func startGithubOauth(rw http.ResponseWriter, r *http.Request) {
	redirect := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(rw, r, redirect, http.StatusSeeOther)
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(rw, "state is incorrect", http.StatusBadRequest)
		return
	}

	tk, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(rw, "couldn't log in", http.StatusInternalServerError)
		return
	}

	ts := githubOauthConfig.TokenSource(r.Context(), tk)

	client := oauth2.NewClient(r.Context(), ts)
}
