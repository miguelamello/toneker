package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/token", getToken)
	http.ListenAndServe(":8000", nil)
}

func getDocumentation(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getToken(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{ token: 1234567890 }")
}
