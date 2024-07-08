package main

import (
	"fmt"
	"html/template"
	"net/http"

	"go.mod/funcs"
)

type AsciiArtData struct {
	InputText string
	AsciiArt  string
}

func Page(w http.ResponseWriter, r *http.Request) {
	if r.Method != "Get" && r.URL.Path == "/" {
		inputText := r.FormValue("input-square")
		data := AsciiArtData{
			InputText: inputText,
		}
		renderTemplate(w, "./html/template.html", data)
	}
	if r.URL.Path != "/asciiart" && r.URL.Path != "/" {
		renderTemplate(w, "./html/404.html", nil)
	}
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inputText := r.FormValue("input-square")
		choice := r.FormValue("choice")
		// fmt.Println(choice)

		if charNotFound(inputText) {
			renderTemplate(w, "./html/ChrNotFound.html", nil)
		} else {

			asciiArt, _ := funcs.Printfinale(inputText, choice)

			data := AsciiArtData{
				InputText: inputText,
				AsciiArt:  asciiArt,
			}
			renderTemplate(w, "./html/template.html", data)
		}
	} else {
		renderTemplate(w, "./html/badRequest.html", nil)
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	t, err := template.ParseFiles(name)
	if name == "./html/404.html" || name == "./html/ChrNotFound.html" {
		w.WriteHeader(404)
	} else if name == "./html/badRequest.html" {
		w.WriteHeader(400)
	}
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "error on your template.html", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {

		// w.WriteHeader(http.StatusNotFound)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func charNotFound(inputText string) bool {
	var exist bool
	exist = false
	for i := 0; i < len(inputText); i++ {
		V := inputText[i]

		if (V < 32 || V > 126) && (V != '\n' && V != '\r') || inputText == "" {
			exist = true
		}

	}
	return exist
}

func main() {
	http.Handle("/asciiart", http.HandlerFunc(asciiArt))

	http.Handle("/", http.HandlerFunc(Page))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Starting server on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
