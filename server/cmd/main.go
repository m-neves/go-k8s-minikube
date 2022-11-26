package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("hi 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8080", nil)

}
