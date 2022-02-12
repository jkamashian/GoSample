package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WOOOOOOOO")
}

func main() {
	http.HandleFunc("/", index_handler)
	//http.HandleFunc("/test", index_handler)
	http.ListenAndServe(":8080", nil)
}
