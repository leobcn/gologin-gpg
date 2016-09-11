package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/openpgp"
)

//generates a random secretString
var secretString = randSeq(1000)

func generateBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	check(err)
	return b
}

func randSeq(l int) string {
	b := generateBytes(l)
	return base64.URLEncoding.EncodeToString(b)
}

//encrypts the secretString
func encrypt() (string, error) {
	keyringFileBuffer, err := os.Open(*publickeyring)
	check(err)

	defer keyringFileBuffer.Close()
	entitylist, err := openpgp.ReadKeyRing(keyringFileBuffer)
	check(err)

	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entitylist, nil, nil, nil)
	check(err)

	_, err = w.Write([]byte(secretString))
	check(err)

	err = w.Close()
	check(err)

	bytesp, err := ioutil.ReadAll(buf)
	check(err)

	encstr := base64.StdEncoding.EncodeToString(bytesp)

	return encstr, err
}
