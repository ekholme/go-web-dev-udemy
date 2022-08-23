# Servers

- Generally, the terms "router," "server," "request router," "mux," "multiplexer," etc. all mean the same thing. They're just talking about servers.
- Clients make requests; servers respond

## TCP

- http servers just run on top of TCP servers, so it's helpful to understand how TCP works
- simple 3-step process for working with a server:
    - listen
    - connect
    - read or write
- look into some of the docs for bufio.NewScanner() -- and practice with this some more!
- but we'll pretty much never use TCP when building a web app, so I'm actually going to skip over the rest of this

## HTTP

- [Handlers](https://pkg.go.dev/net/http#Handler) are pretty much the most important concept of the `http` package. Handlers are interfaces with a ServeHTTP method.
-A [request](https://pkg.go.dev/net/http#Request) is a struct with some information
- **RESUME AT 30 -- UNDERSTANDING LISTEN AND SERVE**