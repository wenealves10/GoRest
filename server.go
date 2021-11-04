package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", UserHandler)
	log.Println("Server running...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
