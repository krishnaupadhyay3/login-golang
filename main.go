package main

import (
	"fmt"
	"login-signup/views"
	"net/http"
	// "blogapp/views"
)

func main() {
	http.HandleFunc("/", views.IndexPage)
	fs := http.FileServer(http.Dir("./web/static"))
	http.HandleFunc("/login", views.LoginPage)
	http.HandleFunc("/register", views.RegisterPage)

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println(http.ListenAndServe(":4000", nil))
}
