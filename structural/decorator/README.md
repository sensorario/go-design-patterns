# Decorator

The decorator pattern provide a lot of benefits when working with legacy code. In this example I have a legacy code. For example a legacy recipe. A decorator add functionality. The functionality in this example is the method Decorate. In the real world Decorate() is represented by a method that must be decorate.

```go
type LegacyRecipe struct {
	Decorate() (string, error)
}

func (lr *LegacyRecipe) Decorate() (string, error) {
	return "Legacy recipe with the following ingredients:", nil
}
```

All object needed to implement the decorator pattern should implement same interface.

```go
type Decorable interface {
	Decorate() (string, error)
}
```

Because of we are treating with a legacy recipe. All new ingredient must contains a `Decorable` recipe.

```go
type NewIngredient struct {
	recipe Decorable
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
```
