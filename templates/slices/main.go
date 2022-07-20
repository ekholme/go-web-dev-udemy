package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

//to pass in a slice
func init() {
	tpl = template.Must(template.ParseFiles("templates/slices/tpl.html"))
}

func main() {
	ppl := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, ppl)
	if err != nil {
		log.Fatalln(err)
	}
}
