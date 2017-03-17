package main

import (
	"./restful"
	"html/template"
	"log"
	"net/http"
	"path"
	"fmt"
	"path/filepath"
)
func ica05(w http.ResponseWriter, r *http.Request) {
	weather := restful.GetWeather("Kristiansand")
	fp := path.Join("html/index.html")
	tmpl, start := template.ParseFiles(fp)
	if start != nil {
		http.Error(w, start.Error(), http.StatusInternalServerError)
		return
	}
	if start := tmpl.Execute(w, weather); start != nil {
		http.Error(w, start.Error(), http.StatusInternalServerError)
	}
}

func main() {
	p, _ := filepath.Abs(" /")
	fmt.Println("Root path:" + p)
	fmt.Println("Starting server:")
	http.HandleFunc("/", ica05)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}