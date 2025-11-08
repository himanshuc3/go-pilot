package racer

import (
	"net/http"
	"time"
)

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
func Racer(a, b string) (winner string) {

	aDuration := measureResponseTime(a)

	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}
