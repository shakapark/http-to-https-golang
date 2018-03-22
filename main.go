package main

import (
	"fmt"
	"net/http"
	"time"
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

func main() {
	handler := http.NewServeMux()
	started := time.Now()

	handler.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Now().Sub(started)
		if duration.Seconds() > 10 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	})

	handler.HandleFunc("/", redirect)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
