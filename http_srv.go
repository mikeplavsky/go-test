package main
import (
  "fmt"
  "net/http"
  "log"
)

type Hello struct {}

func (Hello) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request){
  fmt.Fprint(w, "Hello, World!\n")
}

func main() {
  var h Hello
  log.Fatal(http.ListenAndServe(":8080", h))
}
