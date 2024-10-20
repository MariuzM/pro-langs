package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/about", handleAbout)

	fmt.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our Go server!")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About us")
}
