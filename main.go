package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func main() {

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Servidor escuchando en http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}

//aqui esta el servidor
//creado y
//el tema de
//las rutas
// este archivo se llama main.go
