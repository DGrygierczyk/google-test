package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/increment", incrementCounter)
	http.ListenAndServe(":8081", nil)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}
