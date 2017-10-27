package main

import(
	 "net/http"
)

func game(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w, r, "game.html")

}

func main(){
	
	http.HandleFunc("/",game)
	http.ListenAndServe(":8080", nil)

}