package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {

    // "/" matches "/*" unless that specific route is handled
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL.Path)
        // serves a file if exists, else 404 page not found
        // if URL.Path[1:] is empty index.html is served as default
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    http.HandleFunc("/noFile", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "No File")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
