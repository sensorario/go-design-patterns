package decorator

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewIngredientCannotBeInstantiateWithotDecorableObject(t *testing.T) {
	ni := &NewIngredient{}
	niResult, err := ni.Decorate()
	if err == nil {
		t.Errorf(
			"Decorator cant decorate "+
				"without legacy object, "+
				"not a string with '%s'",
			niResult,
		)
	}
}

func TestNewIngredientMustReturnDefaultText(t *testing.T) {
	ni := &NewIngredient{&LegacyRecipe{}}
	niResult, _ := ni.Decorate()
	if !strings.Contains(niResult, "new ingredient") {
		t.Errorf(
			"Legacy object must contains 'new ingredient', not '%s'",
			niResult,
		)
	}

	fmt.Println(niResult)
}
