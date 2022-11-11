package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Running server on :3000")
	http.HandleFunc("/", Config)
	http.ListenAndServe(":3000", nil)
}

func Config(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Config, %s", os.Getenv("one-time.credit-debit-card"))

}
