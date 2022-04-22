package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func assertStringEqual(
	t *testing.T,
	result string, expected string,
) {
	if result != expected {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func assertOperationsEqual(
	t *testing.T,
	result *SpyCountdownOperations,
	expected []string,
) {
	if !reflect.DeepEqual(result.Calls, expected) {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func TestGreet(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		result := buffer.String()
		expected := "3\n2\n1\nGo!"

		assertStringEqual(t, result, expected)
	})

	t.Run("sleep before every print", func(t *testing.T) {

		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assertOperationsEqual(t, spySleepPrinter, expected)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 2 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.Sleep,
	}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf(
			"should have slept for %v but slept for %v",
			sleepTime,
			spyTime.durationSlept,
		)
	}
}
