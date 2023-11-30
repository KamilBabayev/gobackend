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

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/users":
		userHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		fmt.Fprint(w, "Not Found")
	}
	// if r.URL.Path == "/" {
	// 	homeHandler(w, r)
	// 	return
	// } else if r.URL.Path == "/contact" {
	// 	contactHandler(w, r)
	// 	return
	// } else if r.URL.Path == "/users" {
	// 	userHandler(w, r)
	// 	return
	// } else {
	// 	fmt.Fprint(w, "Not Found")
	// 	return
	// }
}

func main() {
	http.HandleFunc("/", pathHandler)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
