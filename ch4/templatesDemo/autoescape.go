package main

import (
	"html/template"
	"log"
	"os"
)

func autoescape() {
	// func main() {
	const templ = `<p> A: {{.A}} </p> <p> B: {{.B}} </p>`
	t := template.Must(template.New("escapeExample").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}

	data.A = "<script>alert('you have been pwned')</script>"
	data.B = "<b>This is in bold</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
