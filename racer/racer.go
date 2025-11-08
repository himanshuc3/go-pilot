package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// NOTE:
// 1. CAn't rely on external sources for testing,
// therefore mock libraries for http requests are already present.
// This is why we don't need any external dependencies and prevent
// dep injection.
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)

}

// NOTE:
// 1. Select helps us wait on multiple channels as a way to handle concurrency.
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// NOTE:
// 1. Never make channcels using var, since the zero
// value is nil which will block the main thread forever if sent through
// the channel.
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
