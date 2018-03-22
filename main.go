package main
import (
  "net/http"
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
  http.HandleFunc("/", redirect)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
