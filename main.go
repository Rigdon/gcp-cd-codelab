package main

import (
    "fmt"
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello World: %v", time.Now()))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
