# Behavioral » Memento

## Elements

 - Memento - store originator object
 - Originator - build/create mementos
 - Caretaker - keep the memento

## Description

Every time we create new memento object, its state is taken from the originator
element of the design pattern.

```go
func TestNewMementoTakeOriginatorState(t *testing.T) {
	firstState := originator{}
	firstState.state = State{Description: "Idle"}
	mem := firstState.NewMemento()
	if mem.state.Description != "Idle" {
		t.Error("Expected state was not found")
	}
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}
```
In the following test state is changed to "working" to ensure that new state
will be written.

```go
func TestStoresIdleState(t *testing.T) {
	originator := originator{state: State{"Idle"}}
	idleMemento := originator.NewMemento()
	originator.state.Description = "Working"
	originator.ExtractAndStoreState(idleMemento)
	if originator.state.Description != "Idle" {
		t.Error("Unexpected state found")
	}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}
```
