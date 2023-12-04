package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/users":
		userHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> welcome to my app home page</h1>")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>loading users list...</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Our Contact info ...</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	<li><b>Is there a free version?</b> Yes! We offer a free trial for 30 days on any paid plan</li>
	<li><b>What are your support hours?</b>We have support staff answering emails 24/7, though reponse times may be a bit slower on weekends</li>
	<li><b>How do I contact support?</b>Email us - <a href="mailto:support@mail.com">support@mail.com</a></li>
	</ul>
	`)
}

func userIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	fmt.Fprint(w, "user id is: ", userID)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/users", userHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	r.Get("/app/{userID}", userIDHandler)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
