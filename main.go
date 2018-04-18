package main

import (
	"net/http"
	"regexp"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	matched, _ := regexp.MatchString("[0-9]+.[0-9]+.[0-9]+.[0-9]+", req.Host)
	if matched {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found !"))
		return
	}
	// HSTS is a HTTP header that instructs the browser to change all http:// requests to https://.
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/", redirect)

	handler.HandleFunc("/healthz", healthz)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
