package main 

import (
	"fmt"
	"net/http"
)

func handler(http.ResponseWriter, http.Request) {
	fmt.Fprintf("Hi There, I love you")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
