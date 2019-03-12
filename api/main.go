package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello() {
	fmt.Println("Hello world with mux")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8000", r))
}
