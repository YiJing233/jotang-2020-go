// verson 0.1
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Golang yyds！")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("service port:8001")
	err := http.ListenAndServe(":8001",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

