package main
import (
  "net/http"
)

func redirect(w http.ResponseWriter, req *http.Request) {
    // remove/add not default ports from req.Host
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
