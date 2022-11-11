package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Running server on :3000")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s %s %s!", r.URL.Path[1:], os.Getenv("FIRSTNAME"), os.Getenv("LASTNAME"))
}
