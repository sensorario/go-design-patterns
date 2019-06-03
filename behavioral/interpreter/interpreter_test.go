package main

import (
	"strconv"
	"testing"
)

func TestNodesCantBeNormalString(t *testing.T) {
	sentence := "normal"
	i := interpreter{}
	err := i.of(sentence)
	if err == nil {
		t.Error("Si doveva spaccare")
	}
}

func TestFoo(t *testing.T) {
	sentence := "do something"
	i := interpreter{}
	i.of(sentence)
	if i.numberOfWords() != 2 {
		t.Error("Guarda che ti stai sbagliando")
	}
}

func TestPlusOperatorDetector(t *testing.T) {
	sentence := "2 + 3"
	i := interpreter{}
	i.of(sentence)
	if i.contains("+") != true {
		t.Error("dovrebbe conoscere l'operatore +")
	}
	if i.contains("unknown") != false {
		t.Error("non dovrebbe conoscere un operatore sconosciuto")
	}
}

func TestSplitSentencesInSplice(t *testing.T) {
	sentence := "2 + 3"
	i := interpreter{}
	i.of(sentence)
	expected := []string{"2", "+", "3"}
	for ind, _ := range expected {
		tok := i.tokens()
		if expected[ind] != tok[ind] {
			t.Error("non ci siamo")
		}
	}
}

func TestCountNumberOfOperators(t *testing.T) {
	sentence := "2 + 3"
	i := interpreter{}
	i.of(sentence)
	expected := []string{"2", "+", "3"}
	for ind, _ := range expected {
		tok := i.tokens()
		if expected[ind] != tok[ind] {
			t.Error("non ci siamo")
		}
	}
}

func TestExec(t *testing.T) {
	sentence := "5 + 3"
	i := interpreter{}
	i.of(sentence)
	if i.exec() != 8 {
		t.Error([]string{
			"La somma di 5 con 3",
			"non dovrebbe essere",
			strconv.Itoa(i.exec()),
		})
	}
}

func TestMulOperator(t *testing.T) {
	sentence := "5 * 3"
	i := interpreter{}
	i.of(sentence)
	if i.exec() != 15 {
		t.Error([]string{
			"Multiplication between 5 and 3",
			"shouldnt be",
			strconv.Itoa(i.exec()),
		})
	}
}
