//a very basic approach to templating -- string concatenation
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Easy E"

	//we could also get the name from a command-line argument with
	//name := os.Args[1]
	//and run the script as go run main.go Eric to set 'Eric' as the name var

	tpl := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	//create a new index.html file with the above
	nf, err := os.Create("templates/01/index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
