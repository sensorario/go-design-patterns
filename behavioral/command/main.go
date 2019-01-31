package main

func main() {
	c := CommandInvoker{}
	c.addToQueue(&SomeCommand{"Simone"})
	c.addToQueue(&SomeCommand{"Gentili"})
	c.addToQueue(&SomeSpecialCommand{"sensorario"})
}
