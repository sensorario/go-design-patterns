package factory

import (
	"strings"
	"testing"
)

func TestCreateGetGreetingOfWarehouse(t *testing.T) {
	greeting, err := GetTranslator(Italian)
	if err != nil {
		t.Fatal("A GetGreetingment method of type 'Italian' must exists")
	}

	msg := greeting.GetGreeting()
	if !strings.Contains(msg, "Ciao") {
		t.Error("The italian greeting isn't correct")
	}
}

func TestCreateGetGreetingOfEnglish(t *testing.T) {
	greeting, err := GetTranslator(English)
	if err != nil {
		t.Fatal("A GetGreetingment method of type 'English' must exists")
	}

	msg := greeting.GetGreeting()
	if !strings.Contains(msg, "Hello") {
		t.Error("The english greeting isn't correct")
	}
}
