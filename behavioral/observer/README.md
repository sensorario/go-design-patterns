# Behavioral » Observer

## Elements

 - Observable
 - ConcreteObservable
 - Observer
 - ConcreteObserver

## Description

This pattern uncouple an event from its possible handlers. It is useful to
trigger many actions on same events. And it is useful whenever the number of
action of this event grows.

Key elements of this patterns are the publisher and the subscriber. The
publisher will emit an event. All subscriber that handle that event are called
when the particular event is triggered.

## Implementation

First of all in this pattern we needs a publisher and a subscriber. The
publisher will contain the list of subscribers/observers. The subsciber can be
notified (it expose a method `Subscriber.Notify(string)`.

```go
type Subscriber interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Subscriber
}
```

As we do not care about the observer behavior now we will take care about some
functionalities:

 - AddSubscriber
 - RemoveObserver
 - NotifyObservers

```go
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
```
