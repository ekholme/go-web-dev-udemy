# Routing

- A `mux` is shorthand for a multiplexer, which is just another way to say "router." Basically, if a specific request (defined by path, method, and maybe some other stuff) comes in, our router will determine which code to execute.
- We can make a new mux with `http.NewServeMux()`, e.g.

 `mux := http.NewServeMux()`

 And then we can use that to route

 ```
 //this assumes these types have been defined previously and
 //implement the Handler interface
var d hotdog
var c hotcat

mux := http.NewServerMux()
mux.Handle("/dog/", d)
mux.Handle("/cat", c)
//doing "/dog/" with the trailing / will catch "path/dog/something"
//whereas "/cat" would not catch "path/cat/something"

http.ListenAndServe(":8080", mux)
 ```

 Above isn't the most elegant solution, though.

 A slightly better way:

 ```
package main

import (
    "io"
    "net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "some text")
}


func c(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "some different text")
}

func main() {
    http.HandleFunc("/dog/", d)
    http.HandleFunc("/cat", c)

    http.ListenAndServe(":8080", nil)
}
 ```

 The above takes advantage of the fact that if we don't explicitly create a mux (notice that we're passing `nil` into the call to `http.ListenAndServe`), Go will automatically use the [DefaultServeMux](https://pkg.go.dev/net/http), which saves us from typing some boilerplate. We can also use `HandleFunc` rather than defining explicit types that implement handlers.

 Also note that `http.HandleFunc()` takes in, as arguments, a route and a function that has the signature `func(http.ResponseWriter, *http.Request)`

 **RESUME AT 38 THIRD-PARTY SERVEMUX


### Misc
- **Start putting some of this stuff in `Obsidian` in my own words**