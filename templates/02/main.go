package main

import (
	"html/template"
	"log"
	"os"
)

//executing a single template
// func main() {
// 	tpl, err := template.ParseFiles("templates/02/tpl.html")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = tpl.Execute(os.Stdout, nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

//multiple templates
// func main() {
// 	tpl, err := template.ParseFiles("templates/02/tpl.html", "templates/02/tpl2.html")

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	//execute template 1
// 	err = tpl.ExecuteTemplate(os.Stdout, "tpl.html", nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	//execute template 2
// 	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.html", nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// }

//initialize the program by parsing all of the templates
//ensures things just get parsed once
//will make the program slightly more performant
var tpl *template.Template

func init() {

	//Must will panic if there's an error reading the template
	tpl = template.Must(template.ParseGlob("templates/02/*.html"))
}

//a better way to do multiple files via ParseGlob
func main() {
	//execute template 1
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//execute template 2
	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
