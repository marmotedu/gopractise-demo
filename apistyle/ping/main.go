package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pong)
	log.Println("Starting http server ...")
	log.Fatal(http.ListenAndServe(":50052", nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
