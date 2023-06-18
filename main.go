package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "<h1>Hello! Welcome to Gin's World!</h1>")
	} else if r.URL.Path == "/about" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = fmt.Fprint(w, "This blog is used to record programming notes, if you have feedback or suggestions, please contact "+"<a href=\"mailto:3267666759@qq.com\">9yen</a>")
	} else {
		_, _ = fmt.Fprint(w, "<h1>404 NOT FOUNT :(</h1>")
	}
}

func main() {
	fmt.Println("Hello Gin!")
	http.HandleFunc("/", handlerFunc)
	_ = http.ListenAndServe(":3000", nil)
}
