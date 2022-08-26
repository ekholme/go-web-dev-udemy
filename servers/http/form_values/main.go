package main

import (
	"html/template"
	"log"
	"net/http"
)

//can be anything -- just kind of a placeholder for now
type hotdog int

//add a ServeHTTP method so that hotdog implements the Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//parsing the form will return a map of keys and values, which the template provides some details on how to present to the user
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "index.html", req.Form)
}

var tpl *template.Template

//the template is doing a lot of the work here to present the values
func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}
