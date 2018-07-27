package main

import (
	"net/http"
)

// var templates = template.Must(template.ParseFiles("index.html"))

func main()  {
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)

	http.ListenAndServe(":3000", nil)
}
