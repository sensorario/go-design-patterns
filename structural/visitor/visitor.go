package main

import (
	"fmt"
)

type Parts interface {
	AcceptVisitor(PartsVisitor)
}

type Message struct {
	Content string
}

func (aLog *Message) AcceptVisitor(visitor PartsVisitor) {
	visitor.Append(aLog)
}

type LogMessage struct {
	parts []Parts
}

func NewLogMessage() *LogMessage {
	aLog := new(LogMessage)
	aLog.parts = []Parts{
		&Message{"[yyyy-mm-dd]"},
		&Message{"INFO"},
	}
	return aLog
}

func (aLog *LogMessage) AcceptVisitor(visitor PartsVisitor) {
	for _, part := range aLog.parts {
		part.AcceptVisitor(visitor)
	}
}

//Interface of the visitor
type PartsVisitor interface {
	Append(message *Message)
}

//Concrete Implementation of the visitor
type GetMessageVisitor struct {
	FullMessage []string
}

func (aLog *GetMessageVisitor) Append(message *Message) {
	aLog.FullMessage = append(aLog.FullMessage, fmt.Sprintf("Visiting the %v message", message.Content))
}

func main() {
	msg := NewLogMessage()
	visitor := new(GetMessageVisitor)
	msg.AcceptVisitor(visitor)
	fmt.Println(
		"The final message is:",
		visitor.FullMessage,
	)
}
