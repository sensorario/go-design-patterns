package main

import "fmt"
import "testing"

type Subscriber interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Subscriber
}

func (p *Publisher) AddSubscriber(o Subscriber) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Publisher) RemoveObserver(o Subscriber) {
	var indexToRemove int
	for i, observer := range p.ObserverList {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	p.ObserverList = append(
		p.ObserverList[:indexToRemove],
		p.ObserverList[indexToRemove+1:]...,
	)
}

func (p *Publisher) NotifyObservers(message string) {
	for _, observer := range p.ObserverList {
		observer.Notify(message)
	}
}

type TestObserver struct {
	ID      int
	Message string
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Obderver %d: message '%s' received \n", p.ID, m)
	p.Message = m
}

func TestSubject(t *testing.T) {

	testObserver1 := &TestObserver{1, ""}
	testObserver2 := &TestObserver{2, ""}
	testObserver3 := &TestObserver{3, ""}
	publisher := Publisher{}

	t.Run("AddSubscriber", func(t *testing.T) {
		publisher.AddSubscriber(testObserver1)
		publisher.AddSubscriber(testObserver2)
		publisher.AddSubscriber(testObserver3)

		if len(publisher.ObserverList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver2)

		if len(publisher.ObserverList) != 2 {
			t.Fail()
		}
	})

	t.Run("Notify", func(t *testing.T) {
		for _, observer := range publisher.ObserverList {
			printObserver, _ := observer.(*TestObserver)
			message := "hello"
			publisher.NotifyObservers(message)

			if printObserver.Message == "" {
				t.Error()
			}

			if printObserver.Message != message {
				t.Error()
			}
		}
	})

}
