package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	w.Write([]byte("<h1>Hello Full Cycle</h1>"))
	fmt.Fprintf(w, "Hi, im %s. im %s years old", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("techs/techs.txt")

	if err != nil {
		log.Fatalf("Error lendo o arquivo", err)
	}

	fmt.Fprintf(w, "Techs: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s. Pass: %s", user, password)
}