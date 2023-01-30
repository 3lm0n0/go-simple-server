package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST resquest successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
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
	http.HandleFunc("/home", homeHandler)
	// server
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
