// version task1
package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		t, err := template.ParseFiles("input.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func OutputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t, err := template.ParseFiles("output.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, map[string]interface{}{
			"username": r.FormValue("username"),
			"password": r.FormValue("password"),
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/input", handler)
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/output", OutputHandler)
	fmt.Println("service port:8001")
	err := http.ListenAndServe(":8001",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
