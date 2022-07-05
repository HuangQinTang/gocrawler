package main

import (
	"crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler("frontend/view/template.html"))
	http.Handle("/", http.FileServer(http.Dir("frontend/view")))

	err := http.ListenAndServe(":10005", nil)
	if err != nil {
		panic(err)
	}
}
