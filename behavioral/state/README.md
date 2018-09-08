# Behavioral » State

## Description

State is used to encapsulate different states behaviors of the same object.
This pattern is really close to final state machine concepts but it is not its
implementation. State can be seen as a strategy pattern when strategy change on
object properties changes.

## Implementation

In the following example I'll show you state pattern at work to implement a
guess number game.

Main important thin here is the following interface that represents a step of
this game. In terms of state pattern represent a state. A state must execute
the game context and show its name. We will see that a step may chose next step
and decide if the state machine is finished or not. In this implementation a
step returns a boolean value.

```go
type GameStep interface {
	Exec(k *Game) bool
	Name() string
}
```

In main function we have a loop that ends only when Exec function returns
false. Each step have the power to go on with next step o exit the loop.

```go
	for t.CurrentStep.Exec(&t) {
		t.prntState()
	}
```

Before main function we have to analyze Game struct. This struct have the
responsibility to store current step that must be executed. As we said before,
each step can and will change Game's current step. The loop in main function
execute current step.

```go
type Game struct {
	CurrentStep GameStep
	Number       int
	Exit         bool
}

func (a *Game) prntState() {
	fmt.Println(">>", a.CurrentStep.Name())
	if a.CurrentStep.Name() == "end" {
		fmt.Println("")
	}
}
```

Let's see first step. As we are going to implement a guess number game, we have
to implement the Ask state. This state will ask end user to guess the magic
number. How to generate random number is responsibility of main function. In
this state we have to pay attention in Exec method. Here we read from standard
input the value inserted by end user.

In this implementation current step say to the Game if must exit or not. Is a
non necessary step but it is important to see how it is easy to add another
step the queue of the state.

```go
type Ask struct{}

func (s *Ask) Exec(k *AContext) bool {
	var n int
	fmt.Print(">> ")
	fmt.Fscanf(os.Stdin, "%v", &n)
	if n == k.Number {
		k.Exit = true
	} else {
		if n > k.Number {
			fmt.Println(">> you number is greater")
		} else {
			fmt.Println(">> you number is lower")
		}
	}
	k.CurrentState = &CheckState{}
	return true
}

func (s *Ask) Name() string {
	return "guess the number ... "
}

```

```go
func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)

	t := Game{
		CurrentStep: &Ask{},
		Number:       number,
	}

	t.prntState()

	for t.CurrentStep.Exec(&t) {
		t.prntState()
	}
}
```

Finally, you can see other state of current implementation.

```go
type FinishState struct{}

func (s *FinishState) Exec(k *Game) bool {
	fmt.Println("FINISHED !!!")
	return false
}
func (s *FinishState) Name() string {
	return "finish"
}
```

```go
type CheckState struct{}

func (s *CheckState) Exec(k *Game) bool {
	if k.Exit == true {
		k.CurrentStep = &FinishState{}
		return true
	}

	k.CurrentStep = &Ask{}
	return true
}

func (s *CheckState) Name() string {
	return "end"
}
```
