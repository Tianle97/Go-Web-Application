//Tianle Shu
//Question1: GuessingGame

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","game/html")

	fmt.Fprintln(w, "Gussing Game")
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
