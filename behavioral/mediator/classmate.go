package main

type ClassMate struct {
	name        string
	lastMessage string
	forum       string
}

func (a *ClassMate) Class() string {
	return a.forum
}

func (a *ClassMate) Learn(message string) {
	a.lastMessage = message
}

func (a *ClassMate) Learned() string {
	return a.lastMessage
}
