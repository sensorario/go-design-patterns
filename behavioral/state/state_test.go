package main

import "testing"
import "math/rand"
import "time"

func TestStart(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)

	c := AContext{
		CurrentState: &FinishState{},
		Number:       number,
	}

	c.prntState()

	if c.CurrentState.Name() != "finish" {
		t.Error("At FinishState game should be ended")
	}
}
