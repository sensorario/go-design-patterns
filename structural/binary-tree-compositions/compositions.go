package main

import (
	"fmt"
)

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}