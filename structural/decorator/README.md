# Decorator

## Elements

 - Component
 - Decorator
 - Concrete Component
 - Concrete Decorator

## Description

The decorator pattern provide a lot of benefits when working with legacy code. In this example legacy code is represented by a legacy recipe. A decorator add functionality and in this example decorators are new ingredients.

```go
type LegacyRecipe struct {
}
```
Starting from legacy code, we must have a default behavior. For example LegacyRecipe could have a method called `Decorate` that provide the recipe.

```go
func (lr *LegacyRecipe) Decorate() (string, error) {
	return "Original behavior: ", nil
}
```

All object needed to implement the decorator pattern should implement same interface. Because of we must improve the functionality of `Decorate` method we'll create an interface like this.

```go
type Decorable interface {
	Decorate() (string, error)
}
```

And because of we are treating with a legacy recipe, all new ingredient must implement a `Decorable` interface.

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
	return fmt.Sprintf("%s %s,", s, "with decoration"), nil
}
```
As you can see we can use this decorator with:

```go
dec := &NewIngredient{&LegacyRecipe{}}
dev.Decorate() // Original behavior: with decoration
```
