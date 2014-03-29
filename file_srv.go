package main

import (
  "log"
  "net/http"
)

func main() {

  log.Printf("Running...")
  
  log.Fatal(http.ListenAndServe(
    ":8080",
    http.FileServer(http.Dir("."))))

}
