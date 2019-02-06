# Visitor

The visitor pattern is a way to separate algorithm from an object structure. It
is one way to implement the open/closed principle of SOLID. It allows to add
functionalities without object modification. A visitor take the object instance
as input and implements the algorithm.

## Implementation

In the following implementation we will create a visitor that will visit all
cards inside a project. Each card represents a Task or a Bug. Each card
contains a title and some points. Here, card is an interface that for
simplicity implement just a part of a complete card.

Let's start from the card interface. Each card will contains, at least,
GetTitle() and GetPoints() method. In a real world example each card could
contains more and more specific methods. This is just a trivial example.

```go
type Card interface {
	GetTitle() string
	GetPoints() int
}
```

Now let's implement both task and bug object within interface Card.

```go
type Task struct {
	Title string
	Time  int
}

func (b *Task) GetTitle() string {
	return b.Title
}

func (b *Task) GetPoints() int {
	return b.Time
}

type Bug struct {
	Title string
	Time  int
}

func (b *Bug) GetTitle() string {
	return b.Title
}

func (b *Bug) GetPoints() int {
	return b.Time
}
```

Until now we have not yet visitor pattern elements. We just have bugs and
tasks.  In the visitor pattern there always be `Visitable` and `Visitor` items.
Our visitor will visit a card. All visitable item must accept a visitor.

```go
type Visitable interface {
	Accept(v Visitor)
}

type Visitor interface {
	Visit(t Card)
}
```

Now we want that all cards are visitable.

```go
func (b *Task) Accept(v Visitor) {
	v.Visit(b)
}

func (b *Bug) Accept(v Visitor) {
	v.Visit(b)
}
```

And here we have the visitor: a service that sum each cards points. As we can
see the logic is not in Bug object nor in Task.

```go
type EstimationVisitor struct {
	Sum int
}

func (e *EstimationVisitor) Visit(t Card) {
	e.Sum += t.GetPoints()
}
```

Finally, the main function, where it is visible the Visitor in action.

```go
func main() {
	nextRelease := []Visitable{
		&Task{"Do stuff", 1},
		&Task{"Implement Foo Bar", 5},
		&Bug{"Error 500 on resource /foo/bar", 3},
	}

	storyPoint := new(EstimationVisitor)

	for _, i := range nextRelease {
		i.Accept(storyPoint)
	}

	// "Next release is calulated in 9 story points"
	fmt.Println(
		"Next release is calulated in",
		storyPoint.Sum,
		"story points",
	)
}
```
