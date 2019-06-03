# Behavioral » State

## Elements

 - Step: an interface that defines each step;
 - Contest: the one that know current step;
 - Client: the part that iterate steps;

## Description

State is used to encapsulate different states behaviors of the same object.
This pattern is really close to final state machine concepts but it is not its
implementation. State can be seen as a strategy pattern when strategy change on
object properties changes.

## Implementation

In the following example I'll show you state pattern at work to implement a
guess number game.

Main important thing here is the following interface that represents a step of
guess number game. In terms of state pattern represent a state. A state must
execute the game according to game context. We will see that a step may define
Context's next step and decide if the state machine is finished or not. In this
implementation a step returns a boolean value.

```go
type GameStep interface {
	Exec(k *Game) bool
	Name() string
}
```

In main function we have a loop that will ends only when Exec function will
returns false.

```go
	for t.CurrentStep.Exec(&t) {
		t.prntState()
	}
```

Before main function we have to analyze another struct: Game struct. This
struct have the responsibility to store current step and other variable
inherent of the context. In this implementation we have, for example an exit
variable (and a step that decide to exit previous loop of context's exit value
is true).

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
to implement the Ask state. When this state is executed, it will ask end user
to guess the random number. Here we have to pay attention onto Exec method.
Here we read from standard input the value inserted by end user.

In this implementation current step says to the Game's context struct if must
exit or not. Is a non necessary step but it is important to see how it is easy
to add another step the queue of the state.

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

In the main function we have to generate a random number.

```go
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)
```

Then, the game context is generated.

```go
	t := Game{
		CurrentStep: &Ask{},
		Number:       number,
	}
```

And the game can start.

```go
	t.prntState()
	for t.CurrentStep.Exec(&t) {
		t.prntState()
	}
```

Here the full main function when random number is generated, game created and
game loop started.

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
