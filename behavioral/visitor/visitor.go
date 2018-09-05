package main

import "fmt"

type Card interface {
	GetTitle() string
	GetPoints() int
}

type Visitable interface {
	Accept(v Visitor)
}

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

func (b *Task) Accept(v Visitor) {
	v.Visit(b)
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

func (b *Bug) Accept(v Visitor) {
	v.Visit(b)
}

type Visitor interface {
	Visit(t Card)
}

type EstimationVisitor struct {
	Sum int
}

func (e *EstimationVisitor) Visit(t Card) {
	e.Sum += t.GetPoints()
}

func main() {
	nextRelease := []Visitable{
		&Task{"Do stuff", 1},
		&Task{"Implement Foo Bar", 5},
		&Bug{"Error 500 on resource /foo/bar", 3},
	}

	storyPoints := new(EstimationVisitor)

	for _, i := range nextRelease {
		i.Accept(storyPoints)
	}

	fmt.Println(
		"Next release is calulated in",
		storyPoints.Sum,
		"story points",
	)
}
