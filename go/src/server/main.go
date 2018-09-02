package main

import (
	//"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

type config struct {
	RootDirectory string
	DistDirectory string
}

var c config
var configFile string

func rootHandler(w http.ResponseWriter, r *http.Request) {
	file := path.Join(c.DistDirectory, "index.html")
	http.ServeFile(w, r, file)
}
func main() {
	if configFile == "" {
		configFile = "config.json"
	}
	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		log.Println(err)
		log.Fatalln("json decoder failed")
	}
	log.Println(c)
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir(c.DistDirectory))))

	http.Handle("/", r)
	log.Println("Listening on port", 8888, "...")
	http.ListenAndServe(":8888", nil)
}
