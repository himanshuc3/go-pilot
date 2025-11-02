package websites

type WebSiteChecker func(string) bool

// NOTE:
// 1. Structs can have anonymous fields, which is a dope feature but can be abused
type result struct {
	string
	bool
}

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {

		// NOTE:
		// 1. We want the wc to be called in a separate goroutine
		// without having the dependency that wc is itself a goroutine.
		// So, we create an anoymous go routine and call it immediately, iife
		// 2. Also, they maintain lexical scoping, which is helpful in most scenarios.
		// 3. Also, we run into race condition because overriding simultaneous access
		// is not allowed on maps.
		go func() {
			// Writing to a channel
			resultChannel <- result{url, wc(url)}
			// results[url] = wc(url)
		}()
	}

	// NOTE:
	// 1. Each of the goroutines write to the channel, so we need to read from the channel
	// Each write blocks the main thread until the value is read.
	for i := 0; i < len(urls); i++ {
		// Recieve operation blocks until a value is available.
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
