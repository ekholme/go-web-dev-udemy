
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

 **resume at video 11 - composite data structures in templates**