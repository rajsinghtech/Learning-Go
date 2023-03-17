# This program outputs Hello World!
## Explanation 
It's recommended to run `go mod init` in your Go projects to initialize a new module and create a go.mod file which is used to manage dependencies. Without a go.mod file, you won't be able to take advantage of Go's dependency management features and will have to manually manage your project's dependencies.

`go run` is a command used to compile and run a Go program in one step. When you use go run, Go compiles the source code of the program and then immediately executes it, without creating a binary executable file. This means that you don't have to manually compile your Go code before running it, making it a convenient way to quickly run your Go programs during development.

We import the `fmt` package in Go to use the `Println` command because `Println` is a function provided by the `fmt` package. `Println` is a formatting function that prints the specified arguments to the standard output stream and appends a newline character at the end.

## Intialize Project
```
go mod init Hello-World
```
Run main.go 
```
go run main.go
```

[fmt library](https://pkg.go.dev/fmt)
