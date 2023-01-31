package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func formHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method not allowed", http.StatusNotFound)
		return
	}
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	template, err := template.ParseFiles("./static/form.html")
	if err != nil {
		http.Error(res, "500 server error parsing files", http.StatusInternalServerError)
	}
	template.Execute(res, nil)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/home" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method not allowed", http.StatusNotFound)
		return
	}
	template, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(res, "500 server error parsing files", http.StatusInternalServerError)
	}
	template.Execute(res, nil)
}

func main() {
	// Routes
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/form", formHandler)
	// server
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
