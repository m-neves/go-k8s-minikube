package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var url string

func main() {

	go watch()
	go hit()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	//Block
	<-sig
}

func hit() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C

		r, err := http.Get(url)
		if err != nil {
			log.Println("failed to hit target at:", url, "error:", err.Error())
			continue
		}

		log.Println("Got answer from", url)

		b, err := ioutil.ReadAll(r.Body)
		log.Println(string(b))
	}
}

func watch() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		f, err := os.ReadFile("url")
		if err != nil {
			log.Println("failed to open file:", err.Error())
		}

		if url != string(f) {
			url = string(f)
		}

		<-ticker.C
	}
}
