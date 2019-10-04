package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHomePage)

	log.Fatal(http.ListenAndServe(":8080", r))
}
