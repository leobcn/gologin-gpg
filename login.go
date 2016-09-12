package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	secretString = randSeq(1000) //generates new secretString each time the page is loaded
	hex, err := encrypt()        //encrypts the secretString
	check(err)                   //check for errors during encryption

	title := "Login"
	p := &Page{Title: title, GPG: hex} //changes {{Title}} and {{GPG}} values in template
	t, err := template.ParseFiles(*tmpl)
	check(err)

	t.Execute(w, p)
}

//processes login information and informs the server if login fails or succeeds
func processlogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//if GET method redirects back to login page
		http.Redirect(w, r, "/", 301)

	} else {
		//if method is POST try to process login
		r.ParseForm() //parses input

		//Compares username and secret string
		if r.FormValue("username") == *username && r.FormValue("password") == secretString {
			fmt.Print("Login succesful\n")
		} else {
			fmt.Print("Login failed\n")
		}
	}
	//redirects back to login page regardless of login status
	http.Redirect(w, r, "/", 301)
}
