# Serve Files

- The way we want to serve files typically (using standard library) is with `http.FileServer()`. This takes a directory (where your files live) as an argument and works like the following:

```
package main

import (
    "io"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/dog", dog)
    http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    io.WriteString(w, `<img src="toby.jpg">`)
}
```
In the above, assuming we have a file named `toby.jpg` in the root of our directory, that will get served when we navigate to whatever.com/dog. If we just go to whatever.com, we'll see an index that lists all of the files available in the directory, which is kinda cool.

Also note that `http.FileServer()` returns a Handler!
