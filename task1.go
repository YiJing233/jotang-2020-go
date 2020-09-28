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

// handler函数，接收responseWriter结构和http里的Request结构，功能是处理收到的请求
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
	// handleFunc函数，起到路由效果，确定该路径下应该使用哪个处理函数
	http.HandleFunc("/", handler)
	fmt.Println("service port:8001")
	// err为设置监听端口返回的结果，如果失败会有内容
	err := http.ListenAndServe(":8001",nil)
	if err != nil {
		// 本行是print在命令行中的
		log.Fatal("ListenAndServe: ", err)
	}
}

