package main

import (
	"net/http"
	"github.com/wingyplus/gofizzbuzz"
	"log"
)

func main() {
	http.HandleFunc("/say", gofizzbuzz.FizzBuzzHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
