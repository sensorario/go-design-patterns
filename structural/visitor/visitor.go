package main

import (
	"fmt"
)

type Parts interface {
	AcceptVisitor(Visitor)
}

type Visitor interface {
	Visit(message *Message)
}

type Message struct {
	Content string
}

type LogMessage struct {
	parts []Parts
}

type ConcreteVisitor struct {
	FullMessage []string
}

func (aLog *Message) AcceptVisitor(visitor Visitor) {
	visitor.Visit(aLog)
}

func (aLog *LogMessage) AcceptVisitor(visitor Visitor) {
	for _, part := range aLog.parts {
		part.AcceptVisitor(visitor)
	}
}

func (aLog *ConcreteVisitor) Visit(message *Message) {
	aLog.FullMessage = append(
		aLog.FullMessage,
		message.Content,
	)
}

func NewInfoLogString(message *Message) *LogMessage {
	aLog := new(LogMessage)
	aLog.parts = []Parts{
		&Message{"[yyyy-mm-dd]"},
		&Message{"INFO"},
		message,
	}
	return aLog
}

func main() {
	msg := NewInfoLogString(&Message{"Messaggio"})
	visitor := new(ConcreteVisitor)
	msg.AcceptVisitor(visitor)
	fmt.Println(visitor.FullMessage)
}
