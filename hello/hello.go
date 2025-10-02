package main

import "fmt"

const greetPrefix = "Hello, "

// NOTE
// 1. `pkgsite` is used by go docs internally for generating documentation for packages
func hello(name string) string {
	if name == "" {
		name = "stranger"
	}
	return fmt.Sprintf("%s%s!", greetPrefix, name)
}

// NOTE
// 1. Go modules FTW, no more GOPATH headaches!
// 2. `go help` is supposedly helpful!
// 3. Linting using `golangci-lint` is recommended by https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go
// 4. `go fmt` is the default opinionated code formatter used by the Go community.
// 5. Philosophy for testing: Separating "domain logic" from "side effects" (e.g., I/O, network, database, etc.) makes unit testing easier.
// 6. Run `go doc fmt` to check man-pages like documentation
// 7. `go get` - to manage deps vs `go install` - to install a package executable
func main() {
	fmt.Println(hello("you weirdo!"))
}
