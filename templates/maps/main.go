package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

//to pass in a map
func init() {
	tpl = template.Must(template.ParseFiles("templates/maps/tpl.html"))
}

func main() {
	ppl := map[string]string{
		"India":        "Gandhi",
		"America":      "MLK",
		"Also_India":   "Buddha",
		"MidEast":      "Jesus",
		"Also_MidEast": "Muhammed",
		//remember that keys in a map must be unique
	}

	err := tpl.Execute(os.Stdout, ppl)
	if err != nil {
		log.Fatalln(err)
	}
}
