package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "encoding/json"
    "sync"
    "strings"
)
// gorilla mux is a third party router package

type Book struct {
    Id string
    Name string
    Price float64
    Author string
}

type Books []Book

var books Books

var mutex = &sync.Mutex{}

// init will be called before the main runs
func init() {
    books = []Book{
        { "1234", "Something", 23.49, "Someone" },
        { "1235", "Nothing", 35.98, "Nothing" },
    }
}

// return all books on "/" GET request
func returnAllBooks(w http.ResponseWriter, r *http.Request) {
    fmt.Println(books)
    json.NewEncoder(w).Encode(books)
}

// add new book to the books array on "/newBook" POST request
func addNewBook(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()

    book := Book{}
    json.NewDecoder(r.Body).Decode(&book)
    books = append(books, book)
    fmt.Println(books)
    fmt.Fprintf(w, "New book successfully added")

    mutex.Unlock()
}

// update the book on "/book/{id}" PUT request
func updateBookById(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()

    vars := mux.Vars(r)
    key := vars["id"]
    book := Book{}
    json.NewDecoder(r.Body).Decode(&book)

    for i := 0; i < len(books); i++ {
        if strings.Compare(key, books[i].Id) == 0 {
            books[i] = book
        }
    }
    fmt.Println(books)
    fmt.Fprintf(w, "Book with id %s is successly updated", key)

    mutex.Unlock()
}

// delete the book according to its ID on "/book/{id}" DELETE request
func deleteBookById(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()

    vars := mux.Vars(r)
    key := vars["id"]
    i, l := 0, len(books)

    for i = 0; i < l; i++ {
        if strings.Compare(key, books[i].Id) == 0 {
            books = append(books[:i], books[i+1:]...)
            break
        }
    }
    
    if i == l {
        fmt.Fprintf(w, "No books with id %s exist", key)
    } else {
        fmt.Println(books)
        fmt.Fprintf(w, "Book with id %s successfully deleted", key)
    }

    mutex.Unlock()   
}

func main() {

    // create a new router
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", returnAllBooks).Methods("GET")
    myRouter.HandleFunc("/newBook", addNewBook).Methods("POST")
    myRouter.HandleFunc("/book/{id}", updateBookById).Methods("PUT")
    myRouter.HandleFunc("/book/{id}", deleteBookById).Methods("DELETE")

    http.Handle("/", myRouter)

    log.Fatal(http.ListenAndServe(":8080", nil))

}
