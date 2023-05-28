package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	defaultHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
	}

	http.HandleFunc("/", defaultHandler)

	log.Println("Listening for requests at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
