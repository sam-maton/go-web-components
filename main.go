package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/components/", http.StripPrefix("/components/", http.FileServer(http.Dir("components/dist"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	})

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
