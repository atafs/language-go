package main

import (
	"fmt"
	"net/http"
)

//handlers
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)

	fmt.Println("Test")

}
