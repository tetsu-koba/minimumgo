package main

// This should be imported at first
import _ "github.com/tetsu-koba/minimumgo"

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("/"))))
}
