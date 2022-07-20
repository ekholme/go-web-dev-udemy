
# Text Templates
 - 'parsing' a file means reading it into a program
 - file extension doesn't really matter when reading into go (for text templates)
 - a pointer to a template is a "container" into which all templates get parsed and are held. Todd's words -- obv not technical
 - see a very basic example of parsing and printing a text template in `templates/02/...`
 - sort of misc, but something like (filename ...string) as a function arg is a *variadic parameter*, meaning it can take 1 or more arguments

 - a pointer to a template (e.g. *template.Template) will have a bunch of methods available to it, e.g. Execute()

 - see `templates/02/main.go` for parsing multipel templates via ParseGlob

 ## Passing Data into Templates
 - *see `templates/03/...` for putting data into templates
 - {{.}} in a template represents the current piece of data
 - note that you can only pass in one piece of data into a template, but this piece of data can be a struct or something that has lots of things in it
 - see `tpl2.html` for how to use a variable in a template. Not sure why you'd want this, though?

 ## Passing Composite Data into Templates
 - sort of misc, but Go is about types. You create types, attach methods to types, etc.
 - when passing a slice in, we can use {{range .}} to range over the slice, then we can print elements of the slice using {{.}}. We end the range with {{end}}. *See `templates/slices/...`
 - we can see how to pass in maps, structs, and slices of structs in their respective folders
 - it's also fairly straightforward to pass in a struct of slices of structs as well. I don't have the code in these folders, but the syntax is {{range .`content_name`}}
    - see min ~9 in the "passing composite data structures to templates" video

## Passing Functions into Templates
- generally we don't want to put funcs in our templates that, like, pull from the db. But we might want to use some functions that slightly modify the data if necessary.

### Pipelining between functions
- We can pass data from one function to another within a template via piplining. E.g.

{{. | fdbl | fsq}}

Would pass our data (.) to the `fdbl` function and then the result of that function to the `fsq` function
