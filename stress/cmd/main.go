package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var GOROUTINES_COUNT int = 100

func main() {

	c := make(chan bool)

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		start(c)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		stop(c)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/count/", setGoRoutinesCount)

	http.ListenAndServe(":8080", nil)

}

func start(c chan bool) {
	for i := 0; i < GOROUTINES_COUNT; i++ {
		go loop(c)
	}
}

func stop(c chan bool) {
	for i := 0; i < GOROUTINES_COUNT; i++ {
		c <- true
	}
}

func loop(c chan bool) {

loop:
	for {
		fmt.Println("looping")

		select {
		case <-c:
			fmt.Println("finished")
			break loop
		default:
		}
	}

}

func setGoRoutinesCount(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	cnt, err := strconv.Atoi(path[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Setting GOROUTINES_COUNT to", cnt)

	GOROUTINES_COUNT = cnt
}
