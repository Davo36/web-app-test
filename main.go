// Test Bootstrap in small project

package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

var templateSrc = "./html/"

func parseTemplate(name string) *template.Template {
	t := template.New(name)
	t, _ = t.ParseFiles(templateSrc + name)
	return t
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := Page{"Death to All Rats!", "And possums and stoats and feral cats and ..."}
	t := parseTemplate("index.html")

	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)

	cssFs := http.FileServer(http.Dir("C:/Users/OEM/Google Drive/Programming/go/src/web-app-test/css"))
	http.Handle("/css", cssFs)
	// jsFs := http.FileServer(http.Dir("/home/user/www/assets/js"))
	// http.Handle("/js", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
