package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "Hello auto-pack."
	}
	io.WriteString(w, greeting)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
