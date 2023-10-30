package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	SetupRoutes()

	fmt.Printf("Serveur démarré sur le port : http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func SetupRoutes() {
	http.Handle("/", handlers.Middleware(http.HandlerFunc(handlers.GetHandler)))
	http.Handle("/ascii-art", handlers.Middleware(http.HandlerFunc(handlers.AsciiArtHandler)))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/error404", func(w http.ResponseWriter, r *http.Request) {
		handlers.ErrorPageHandler(w, r, http.StatusNotFound, "error404")
	})
	http.HandleFunc("/error400", func(w http.ResponseWriter, r *http.Request) {
		handlers.ErrorPageHandler(w, r, http.StatusBadRequest, "error400")
	})
	http.HandleFunc("/error500", func(w http.ResponseWriter, r *http.Request) {
		handlers.ErrorPageHandler(w, r, http.StatusInternalServerError, "error500")
	})
}
