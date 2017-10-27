package main

import (
	"net/http"
	"html/template"
	"log"
)

type Message struct{

	Message string

}

func game(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w, r, "guess.html")

}

func guess(w http.ResponseWriter, r *http.Request){

	mess := &Message{
		Message : "Guess a number from 1 to 20."}
	
	t, e := template.ParseFiles("guess.tmpl")
	
	if e != nil{
		log.Print("error",e)
	}

	e = t.Execute(w,mess)
	if e != nil{
		log.Print("error",e)
	}
}

func main(){

	http.HandleFunc("/",game)
	http.HandleFunc("/guess",guess)
	http.ListenAndServe(":8080",nil)

}