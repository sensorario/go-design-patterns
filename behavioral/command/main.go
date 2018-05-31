package main

import "fmt"

type Command interface {
	Execute()
}

type SomeCommand struct {
	message string
}

func (s *SomeSpecialCommand) Execute() {
	fmt.Println(" >>> ", s.message)
}

type SomeSpecialCommand struct {
	message string
}

func (s *SomeCommand) Execute() {
	fmt.Println(s.message)
}

func PrintlnCommand(s string) Command {
	fmt.Println("Creating Command")
	return &SomeCommand{
		message: s,
	}
}

func SpecialPrintlnCommand(s string) Command {
	fmt.Println("Creating Command")
	return &SomeSpecialCommand{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (c *CommandQueue) processQueue() {
	for i := range c.queue {
		c.queue[i].Execute()
	}
}

func (c *CommandQueue) addToQueue(i Command) {
	c.queue = append(c.queue, i)
	if len(c.queue) == 3 {
		c.processQueue()
		c.queue = []Command{}
	}
}

func main() {
	c := CommandQueue{}
	c.addToQueue(PrintlnCommand("Ciaone"))
	c.addToQueue(PrintlnCommand("Bar"))
	c.addToQueue(SpecialPrintlnCommand("Foo"))
}
