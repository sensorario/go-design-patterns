package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	cs := ConsoleStrategy{}
	output := cs.BuildOutput()

	if output != "ConsoleStrategy" {
		t.Error("wrong output message")
	}
}
