package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Page struct {
	Valeur string
}

const port = 1705

var Templateshome = template.Must(template.ParseFiles("./static/templates/home.html"))
var Templateshang = template.Must(template.ParseFiles("./static/templates/hang.html"))
var Templatesmore = template.Must(template.ParseFiles("./static/templates/more.html"))

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	StartGame("./words.txt")
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/hang", hangHandler)
	http.HandleFunc("/more", moreHandler)
	http.Handle("/", http.FileServer(http.Dir("./static/style")))
	fmt.Println("http://localhost:1705/home - server started on port", port)
	http.ListenAndServe(":1705", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Valeur: "Home page"}
	err := Templateshome.ExecuteTemplate(w, "home.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hangHandler(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("answer") //get value of form
	print(val)
	FindAndReplace(val)
	testEndGame()
	if r.Method != http.MethodPost {
		StartGame("./words.txt")
	}
	Templateshang.Execute(w, hangman)
}

func moreHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Valeur: "More page"}
	err := Templatesmore.ExecuteTemplate(w, "more.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
