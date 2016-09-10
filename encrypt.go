package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/openpgp"
)

//generates a random secretString
var secretString = randSeq(1000)

//characters that will be used for a random string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123467890!£$%^&*_+=-?@~#:;}][|")

//function that generates the random string of characters with n number of characters
func randSeq(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
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
