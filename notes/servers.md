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