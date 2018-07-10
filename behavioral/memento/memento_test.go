package main

import (
	"testing"
)

func TestNewMementoTakeOriginatorState(t *testing.T) {
	firstState := originator{}
	firstState.state = State{Description: "Idle"}
	mem := firstState.NewMemento()
	if mem.state.Description != "Idle" {
		t.Error("Expected state was not found")
	}
}

func TestCareTakerKeepAlMementoInsideAList(t *testing.T) {
	firstState := originator{}
	firstState.state = State{Description: "Idle"}
	mem := firstState.NewMemento()
	careTaker := careTaker{}
	lenBeforeAddedMemento := len(careTaker.mementoList)
	careTaker.Add(mem)
	if len(careTaker.mementoList) != lenBeforeAddedMemento+1 {
		t.Errorf("No new elements were added on the list")
	}
}

func TestCareTaker(t *testing.T) {
	originator := originator{}
	careTaker := careTaker{}
	originator.state = State{"foo"}
	careTaker.Add(originator.NewMemento())
	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Fatal(err)
	}
	if mem.state.Description != "foo" {
		t.Error("Unexpected state")
	}
}

func TestCareTakerDetectIfNonExistentIndexIsRequested(t *testing.T) {
	originator := originator{}
	careTaker := careTaker{}
	originator.state = State{"foo"}
	careTaker.Add(originator.NewMemento())
	_, err := careTaker.Memento(0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatal("An error is expected")
	}
}

func TestStoresIdleState(t *testing.T) {
	originator := originator{state: State{"Idle"}}
	idleMemento := originator.NewMemento()
	originator.state.Description = "Working"
	originator.ExtractAndStoreState(idleMemento)
	if originator.state.Description != "Idle" {
		t.Error("Unexpected state found")
	}
}
