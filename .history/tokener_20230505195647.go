package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Tokener okay")
}