package main

import "fmt"

type Observer interface {
	Notifiy(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (p *Publisher) AddObserver(o Observer)        {}
func (p *Publisher) RemoveObserver(o Observer)     {}
func (p *Publisher) NotifyObserver(message string) {}

func main() {
	fmt.Println("vim-go")
}
