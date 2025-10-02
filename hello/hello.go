package main

import "fmt"

const spanish = "Spanish"
const englishGreetPrefix = "Hello, "
const spanishGreetPrefix = "Hola, "
const french = "French"
const frenchGreetPrefix = "Bonjour, "

// NOTE
// 1. `pkgsite` is used by go docs internally for generating documentation for packages
// 2. It seems like golang doesn't have overloading or optional arguments
func hello(name, language string) string {
	prefix := greetingPrefix(language)
	if name == "" {
		name = "stranger"
	}

	return fmt.Sprintf("%s%s!", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		// NOTE
		// 1. `prefix` is a named return value, which is initialized to the zero value of the type ""
		prefix = frenchGreetPrefix
	case spanish:
		prefix = spanishGreetPrefix
	default:
		prefix = englishGreetPrefix
	}
	return
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
	fmt.Println(hello("you weirdo!", ""))
}
