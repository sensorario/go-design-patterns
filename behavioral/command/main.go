package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type SomeCommand struct {
	message string
}

func (s *SomeCommand) Execute() {
	fmt.Println(s.message)
}

type SomeSpecialCommand struct {
	message string
}

func (s *SomeSpecialCommand) Execute() {
	message := "@" + s.message
	fmt.Println(message)
}

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

func main() {
	c := CommandInvoker{}
	c.addToQueue(&SomeCommand{"Simone"})
	c.addToQueue(&SomeCommand{"Gentili"})
	c.addToQueue(&SomeSpecialCommand{"sensorario"})
}
