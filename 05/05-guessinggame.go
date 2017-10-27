//Tianle Shu

package main

import (
	"net/http"
	"html/template"
	"log"
	"strconv"
	"time"
	"math/rand"
)

type message struct{
	
		Message string
	
	}
	
	func game(w http.ResponseWriter, r *http.Request){
	
		http.ServeFile(w, r, "guess.html")
	
	}
	
	func guess(w http.ResponseWriter, r *http.Request){
	
		num := getnum(1,20)
		//set  a new cookie

		if _,err := r.Cookie("target");
		err != nil{
			http.SetCookie(w, &http.Cookie{Name: "target", Value: strconv.Itoa(num)})
		}
         //create a new variable and assign the struct and string value
		mess := &message{
			Message :"Guess a number from 1 to 20"}	
		
			t,err := template.ParseFiles("guess.tmpl")
                 //error handling
			if err != nil{
				log.Print("error",err)
			}
			
			err = t.Execute(w,mess)
			if err != nil{
				log.Print("error",err)			
			}
		}

		func getnum(min,max int)int{
			rand.Seed(time.Now().UnixNano())
			return rand.Intn(max - min) + min
	
		}
	
	func main(){ 
	
		http.HandleFunc("/",game)
		http.HandleFunc("/guess",guess)
		http.ListenAndServe(":8080",nil)
	
	}