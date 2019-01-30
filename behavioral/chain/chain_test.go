package main

import "testing"
import "fmt"

func TestFoo(t *testing.T) {
	chain := TranslatorChain{
		Start: &EnglishTranslator{
			&FrenchTranslator{
				&SpanishTranslator{},
			},
		},
	}

	fmt.Println("vim-go")

	if chain.CountRings() != 3 {
		t.Error("There should be three rings")
	}

	str, _ := chain.Translate("ciao")
	if str != "hola" {
		t.Error("Translation should be kept from spanish ring")
	}

	str, _ = chain.Translate("casa")
	if str != "maison" {
		t.Error("Translation should be kept from french ring")
	}

	str, _ = chain.Translate("topo")
	if str != "mouse" {
		t.Error("Translation should be kept from english ring")
	}

	str, _ = chain.Translate("cocomero")
	if str != "watermelon" {
		t.Error("Translation should be kept from english ring")
	}
}
