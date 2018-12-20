package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type State struct {
	One    string
	Two    string
	Three  string
	Four   string
	Five   string
	Six    string
	Seven  string
	Eight  string
	Nine   string
	Ten    string
	Eleven string
	Twelve string
}

func main() {
	state := State{
		"true", "false", "true", "false",
		"false", "true", "false", "true",
		"true", "false", "false", "true",
	}

	templates := template.Must(template.ParseFiles("templates/ui.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "ui.html", state); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
