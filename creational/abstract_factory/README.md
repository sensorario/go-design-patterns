# Creational » Abstract Factory

## Description

The difference from factory design pattern is that with this pattern it is possible to manage complex data type.

## Implementation

The better way to explain this pattern is to show its power. We need just three steps. First of all we create a factory of italian cars. Second we can ask to this factory a Ferrari.

```go
func TestFoo(t *testing.T) {
	itaF, _ := BuildFactory(ItalianType)
	m, _ := itaF.Build(FerrariModel)

	car, ok := m.(Vehicle)
	if !ok {
		t.Fatal("Invalid model")
	}

	t.Logf("%v car has %d wheels", car.GetModelName(), car.NumOfWheels())
}
```
