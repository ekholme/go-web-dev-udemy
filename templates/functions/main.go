package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

//define a FuncMap that includes the functions we want to send to our template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

//read in the template like this
//note that the funcs need to be defined before you parse the file, which is why we need to read it in like this
func init() {
	tpl = template.Must(template.New("tpl.html").Funcs(fm).ParseFiles("templates/functions/tpl.html"))
}

//define the firstThree func
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "Be cool, brother",
	}

	j := sage{
		Name:  "Jesus",
		Motto: "Be nice to your friends",
	}

	sages := []sage{b, j}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
