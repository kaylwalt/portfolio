package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../dist/index.html")
}
func main() {
	//fileServer := http.FileServer(http.Dir("../dist"))
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("../dist"))))

	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
