package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	spySleeperPrinter := &SpyCountdownOperations{}
	Countdown(spySleeperPrinter, spySleeperPrinter)

	expected := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}
	if !reflect.DeepEqual(spySleeperPrinter.Calls, expected) {
		t.Errorf("wanted calls %v but got %v", expected, spySleeperPrinter.Calls)
	}
}

// NOTE:
// 1. Trying to minimize mocking for testing since it leads to code smell identification.
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	// NOTE:
	// 1. Passing a function which is a method on a struct which takes object's context as well.
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
