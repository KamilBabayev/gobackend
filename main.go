package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> welcome to my app home page</h1>")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>loading users list...</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Our Contact info ...</h1>")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/users", userHandler)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
