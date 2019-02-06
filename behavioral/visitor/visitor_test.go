package main

import "testing"

func TestVisitableBuildItemWithTasksAndBugs(t *testing.T) {
	c := []Visitable{
		&Task{"Do stuff", 1},
		&Task{"Implement Foo Bar", 5},
		&Bug{"Error 500 on resource /foo/bar", 3},
	}

	storyPoints := new(EstimationVisitor)

	for _, i := range c {
		i.Accept(storyPoints)
	}
}
