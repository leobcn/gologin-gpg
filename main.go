package main

import (
	"flag"
	"fmt"
	"net/http"
)

//Page variables to be used within HTML login template
type Page struct {
	Title string
	GPG   string
}

var port, publickeyring, username, tmpl *string

func main() {
	//assign flags
	port = flag.String("port", "9090", "port number")
	publickeyring = flag.String("public_key", "", "GPG public keyring")
	username = flag.String("user", "username", "change username")
	tmpl = flag.String("template", "login.tmpl", "change template location")
	flag.Parse()

	//handle http requsts
	http.HandleFunc("/", login)
	http.HandleFunc("/processlogin", processlogin)
	err := http.ListenAndServe(":"+*port, nil)
	check(err)

}

//error handling function
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
