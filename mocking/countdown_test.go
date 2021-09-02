package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("counts down from 3 and says go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
go`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before each print", func(t *testing.T) {
		countdownOperationsSpy := &CountdownOperationsSpy{}

		Countdown(countdownOperationsSpy, countdownOperationsSpy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, countdownOperationsSpy.Calls) {
			t.Errorf("wanted calls %v, got %v", want, countdownOperationsSpy.Calls)
		}
	})
}
func TestConfigurableSleeper(t *testing.T) {
	sleeptime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleeptime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleeptime {
		t.Errorf("should have slept for %v but slept for %v", sleeptime, spyTime.durationSlept)
	}
}
