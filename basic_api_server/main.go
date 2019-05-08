package main

import (
    "log"
    "net/http"
    "encoding/json"
)

type Book struct {
    Name string `json:"name"`
    Price float64 `json:"price"`
    Author string `json:"author"`
}

type Books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {

        books := Books{
            { Name: "Harry Potter", Price: 49.99, Author: "J.K. Rowling", }, 
            { Name: "Someone", Price: 99.99, Author: "Someone", },
        }
        w.Header().Set("Content-Type", "application/json")
        //w.WriteHeader(http.StatusCreated)

        b, err := json.Marshal(books)

        if err != nil {
            log.Println(err)
        }

        w.Write(b)

        // json.NewEncoder().Encode(books) will also do
}

func main() {

    http.HandleFunc("/getBooks", getBooks)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
