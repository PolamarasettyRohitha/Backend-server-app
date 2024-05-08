package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/home", userAuth(home))
	mux.HandleFunc("/login", login)
	mux.HandleFunc("POST /addBook", adminAuth(addBook))
	mux.HandleFunc("POST /deleteBook", adminAuth(deleteBook))

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
