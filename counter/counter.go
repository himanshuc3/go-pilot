package counter

import "sync"

// NOTE:
// 1. Use channels when passing ownership of data
// 2. Use mutexes for managing state
// 3. `go vet` is similar to eslint for golang.
type Counter struct {
	value int
	// NOTE:
	// 1. It can be an anonymous embedding (sync.Mutex without a name) into the struct
	// which is an implicit composition and better API ig.
	mu sync.Mutex
}

// NOTE:
// 1. Constructor function which is better for signalling
// how we want the structures to be used/initialized.
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
