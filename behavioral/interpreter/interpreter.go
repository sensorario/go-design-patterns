package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type interpreter struct {
	sentence string
}

func (i *interpreter) tokens() []string {
	return strings.Split(i.sentence, " ")
}

func (i *interpreter) exec() int {
	sum := 0
	tokens := i.tokens()
	for k, item := range tokens {
		if item == "*" {
			fmt.Println(i.tokens())
			a, _ := strconv.Atoi(string(tokens[k-1]))
			b, _ := strconv.Atoi(string(tokens[k+1]))
			return a * b
		}

		if item != "+" {
			number, _ := strconv.Atoi(item)
			sum += number
		}
	}
	return sum
}

func dict() map[string]string {
	var m map[string]string
	m = make(map[string]string)
	m["+"] = "plus"
	return m
}

func (i *interpreter) contains(s string) bool {
	m := dict()
	if _, ok := m[s]; ok {
		return true
	}

	return false
}

func (i *interpreter) of(s string) error {
	if s == "normal" {
		return errors.New("non va")
	}
	i.sentence = s
	return nil
}

func (i *interpreter) numberOfWords() int {
	s := i.tokens()
	return len(s)
}
