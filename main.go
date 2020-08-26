package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// book stuct

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appllication/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appllication/json")

	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

		json.NewEncoder(w).Encode(&Book{})
	}

}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appllication/json")
	upid := mux.Vars(r)
	for i, val := range books {
		if val.ID == upid["id"] {
			books = append(books[:i], books[i+1:]...)
			var updatebook Book
			_ = json.NewDecoder(r.Body).Decode(&updatebook)
			updatebook.ID = upid["id"]
			books = append(books, updatebook)
			json.NewEncoder(w).Encode(updatebook)
			return
		}
	}

}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appllication/json")

	dele := mux.Vars(r)
	for i, val := range books {
		if val.ID == dele["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}

		json.NewEncoder(w).Encode(books)

	}

}
func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appllication/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func main() {
	//intial router
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "18797928", Title: "man searching for a bird", Author: &Author{Firstname: "Gary", Lastname: "vchunk"}})
	books = append(books, Book{ID: "2", Isbn: "71737939", Title: "go lang", Author: &Author{Firstname: "ashik", Lastname: "patel"}})
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/book", createBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))

}
