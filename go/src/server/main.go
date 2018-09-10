package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

type config struct {
	RootDirectory string
	DistDirectory string
	EnableHTTPS   string
	CertDirectory string
}

var c config
var configFile string
var httpPort = ":80"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	file := path.Join(c.DistDirectory, "index.html")
	http.ServeFile(w, r, file)
}
func makeServerFromMux(mux http.Handler) *http.Server {
	// set timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
}

func makeHTTPServer() *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir(c.DistDirectory))))

	return makeServerFromMux(router)

}

func makeHTTPToHTTPSRedirectServer() *http.Server {
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		newURI := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, newURI, http.StatusFound)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", handleRedirect)
	return makeServerFromMux(router)
}

func main() {
	if configFile == "" {
		log.Println("No config file given, setting to config.json")
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
	var m *autocert.Manager
	if c.EnableHTTPS == "true" {
		var httpsServer *http.Server

		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("www.kaylwalton.com", "kaylwalton.com"),
			Cache:      autocert.DirCache(c.CertDirectory),
		}
		httpsServer = makeHTTPServer()
		httpsServer.Addr = ":443"
		//httpsServer.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		httpsServer.TLSConfig = m.TLSConfig()
		go func() {
			fmt.Printf("Starting HTTPS server on %s\n", httpsServer.Addr)
			err := httpsServer.ListenAndServeTLS("", "")
			if err != nil {
				log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
			}
		}()
	}

	var httpServer *http.Server
	if c.EnableHTTPS == "true" {
		httpServer = makeHTTPToHTTPSRedirectServer()
	} else {
		httpServer = makeHTTPServer()
	}
	if m != nil {
		httpServer.Handler = m.HTTPHandler(httpServer.Handler)
	}

	httpServer.Addr = httpPort
	log.Printf("Starting HTTP server on %s\n", httpPort)
	serverErr := httpServer.ListenAndServe()
	if serverErr != nil {
		log.Fatalf("httpSrv.ListenAndServe() failed with %s", serverErr)
	}
}
