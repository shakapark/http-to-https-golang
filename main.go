package main
import (
  "fmt"
  "net/http"
  "time"
  "regexp"
)

func redirect(w http.ResponseWriter, req *http.Request) {
    matched, _ := regexp.MatchString("[0-9]+.[0-9]+.[0-9]+.[0-9]+", req.Host)
    if matched == true {
      w.WriteHeader(404)
      return
    }
    target := "https://" + req.Host + req.URL.Path
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    http.Redirect(w, req, target,
            http.StatusMovedPermanently)
}

func main() {
  started := time.Now()

  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    duration := time.Now().Sub(started)
    if duration.Seconds() > 10 {
      w.WriteHeader(500)
      w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
    } else {
      w.WriteHeader(200)
     w.Write([]byte("ok"))
    }
  })

  http.HandleFunc("/", redirect)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
