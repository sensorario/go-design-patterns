package decorator

import (
	"errors"
	"fmt"
)

type LegacyRecipe struct {
}

type Decorable interface {
	Decorate() (string, error)
}

type NewIngredient struct {
	recipe Decorable
}

func (lr *LegacyRecipe) Decorate() (string, error) {
	return "Legacy recipe with the following ingredients:", nil
}

func (ni *NewIngredient) Decorate() (string, error) {
	if ni.recipe == nil {
		return "", errors.New("decorable recipe is needed")
	}
	s, err := ni.recipe.Decorate()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "new ingredient"), nil
}
