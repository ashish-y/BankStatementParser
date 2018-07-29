
package main

import (
  "log"
  "net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
