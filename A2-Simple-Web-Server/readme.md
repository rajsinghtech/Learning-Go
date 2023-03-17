# Simple Web Server
This project uses the package net/http which provides HTTP client and server implementations. In this case we utilize the HandleFunc method and ListenAndServe method.

## Explanation
HandleFunc
- `Handle` and `HandleFunc` add handlers to DefaultServeMux.
- First argument being the path.
- Second argument being the handler function to be called.

ListenAndServe
- ListenAndServe starts an HTTP server with a given address and handler. 
- In Go, `nil` is a predeclared identifier that represents a null value or a zero value for a pointer, interface, map, slice, channel or function type. This is the second argument because `http.ListenAndServe` uses the default server configuration if nil is passed as the second argument.
 
Intialize Project
```
go mod init Web-Server
```
Run main.go 
```
go run main.go
```

[fmt library](https://pkg.go.dev/fmt)
[net/http library](https://pkg.go.dev/net/http)
