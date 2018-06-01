# Behavioral » Command

## Description

While strategy pattern is focused on changing algorithm, in Command pattern the
focus is on invocation of something or on the abstraction of some type. A command
pattern is commonly seen as a container.

## Implementation

This pattern requires an interface for all commands that could be executed in a
certain moment.

```go
type Command interface {
	Execute()
}
```

Here first example of command:

```go
type SomeCommand struct {
	message string
}

func (s *SomeCommand) Execute() {
	fmt.Println(s.message)
}
```

And here a second one:


```go
type SomeSpecialCommand struct {
	message string
}

func (s *SomeSpecialCommand) Execute() {
	message := "@" + s.message
	fmt.Println(message)
}
```

In the `SomeCommand` there is a simple `Println` output. Instead, using another
`SomeSpecialCommand` we also add a `@` symbol at the beginning of the string.

Finally we have an invoker, a component of this pattern that:

 - append commands into queue;
 - process entire queue;

```go
type CommandInvoker struct {
	queue []Command
}

func (c *CommandInvoker) processQueue() {
	for i := range c.queue {
		c.queue[i].Execute()
	}
}

func (c *CommandInvoker) addToQueue(i Command) {
	fmt.Println("Appending command")
	c.queue = append(c.queue, i)
	if len(c.queue) == 3 {
		c.processQueue()
		c.queue = []Command{}
	}
}
```
