package main

import (
    "fmt"
    "net/http"
    "log"
    "sync"
    "strconv"
)

// a web server in golang in asynchrounous
// hence we need to gaurd the critical section using 
// mutex to prevent race condition bugs

// create a counter for critical section
var counter int
var mutex = &sync.Mutex{} // always returns the same mutex address

func echoString(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
    // lock the mutex
    mutex.Lock()
    // critical section
    counter++
    fmt.Fprintf(w, strconv.Itoa(counter))
    // unlock the mutex
    mutex.Unlock()
}

func main() {
    
    http.HandleFunc("/", echoString)

    http.HandleFunc("/incrementCounter", incrementCounter)

    log.Fatal(http.ListenAndServe(":8080", nil))

}
