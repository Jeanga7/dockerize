package handlers

import (
	"ascii-art-web/ascii-art"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	ReadyToConvert bool
	AsciiDrawing   string
	Text           string
	Template       string
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderTemplate(w, "error404", nil, http.StatusNotFound)
		return
	}
	data := Data{Text: "Hello World !!!"}
	RenderTemplate(w, "index", data, http.StatusOK)
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		RenderTemplate(w, "error404", nil, http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			RenderTemplate(w, "error500", nil, http.StatusInternalServerError)
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")
		if text == "" || banner == "" {
			RenderTemplate(w, "error400", nil, http.StatusBadRequest)
			return
		}

		if !ascii.IsPrintable(text) || (banner != "standard" && banner != "shadow" && banner != "thinkertoy") {
			RenderTemplate(w, "error400", nil, http.StatusBadRequest)
			return
		}

		lines := ascii.Open_File(banner)
		t := ascii.Finder(text)
		if len(lines) != 855 {
			RenderTemplate(w, "error500", nil, http.StatusInternalServerError)
			return
		}

		asciiDrawing := ascii.Display(lines, t)

		response := &Data{
			ReadyToConvert: true,
			AsciiDrawing:   asciiDrawing,
			Text:           text,
			Template:       banner,
		}

		RenderTemplate(w, "index", response, http.StatusOK)
		return
	}

	RenderTemplate(w, "index", nil, http.StatusOK)
}

func ErrorPageHandler(w http.ResponseWriter, r *http.Request, statusCode int, pageName string) {
	w.WriteHeader(statusCode)
	filePath := fmt.Sprintf("./templates/%s.html", pageName)
	http.ServeFile(w, r, filePath)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware: Before handling request")
		next.ServeHTTP(w, r)
		log.Println("Middleware: After handling request")
	})
}
