package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
}

var key = []byte{}

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("sguessou:return1771")))

	pass := "return1771"
	hashPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	comparePassword(pass, hashPass)
	if err != nil {
		log.Fatalln("Not Logged in!")
	}
	log.Println("Logged in!")

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8087", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Marc",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}
	log.Println("Person:", p1)
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing message: %w", err)
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)
	}
	same := hmac.Equal(newSig, sig)
	return same, nil
}
