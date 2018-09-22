package main

type Student interface {
	Class() string
	Learn(message string)
	Learned() string
}

type Mediator interface {
	TeachesTo(c *Student)
	Spread(messqger string)
}
