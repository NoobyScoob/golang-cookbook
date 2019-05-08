package main

import (
    "fmt"
    "net/http"
    "log"
    "html"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// EscapeString escapes special characters like "<" to become "&lt;"
        // %q double quotes the string as in source
        fmt.Fprintf(w, "Hello - %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/basic", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is BASIC")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
