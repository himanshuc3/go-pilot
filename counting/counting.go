package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const write = "write"
const sleep = "sleep"

// NOTE:
// 1. Implements both the Sleeper and io.Writer interfaces.
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	// NOTE:
	// 1. Returning the default values even though no explicit values are returned
	// because naming it function signature
	return
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// NOTE:
// 1. Spying on our sleeper's activities
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// NOTE:
// 1. A little more convulated than we've implemented till now
// in terms of mixing concrete types and interfaces (to mimic classes).
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

// NOTE:
// 1. Learning more about stubs, mocks and spies by Martin Fowler:
// https://martinfowler.com/articles/mocksArentStubs.html
func Countdown(out io.Writer, sleeper Sleeper) {
	const finalWord = "Go!"
	const coundownStart = 3
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		// NOTE:
		// 1. Mocking time.Sleep so that we can inject our dependency for testing
		// time.Sleep(time.Second * 1)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	// NOTE:
	// 1. Need to pass a pointer to DefaultSleeper since the interface implementation (Sleeper) used a pointer receiver.
	Countdown(os.Stdout, &DefaultSleeper{})
}
