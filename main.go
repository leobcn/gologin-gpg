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

var port, publickeyring, username, tmpl, tlsCert, tlsKey *string
var httpDisable, tls *bool

func main() {
	//assign flags
	port = flag.String("port", "9090", "port number")
	publickeyring = flag.String("publicKey", "pubring.gpg", "GPG public keyring")
	username = flag.String("user", "username", "change username")
	tmpl = flag.String("template", "login.tmpl", "change template location")
	tls = flag.Bool("tls", false, "use TLS")
	tlsCert = flag.String("cert", "cert.pem", "location of cert.pem file")
	tlsKey = flag.String("key", "key.pem", "location of key.pem file")
	flag.Parse()

	//handle http requests
	http.HandleFunc("/", login)
	http.HandleFunc("/processlogin", processlogin)

	if *tls == true {
		err := http.ListenAndServeTLS(":"+*port, *tlsCert, *tlsKey, nil)
		check(err)
	} else if *tls == false {
		err := http.ListenAndServe(":"+*port, nil)
		check(err)
	}

}

//error handling function
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
