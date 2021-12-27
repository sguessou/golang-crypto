package main

import (
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

func main() {

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
