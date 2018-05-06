# Creational » Abstract Factory

## Intent

 - provide an interface for creating families of related or dependent objects
 - a hierarchy that encapsulates: many possible "platforms", and the construction of a suite of "products".
 - the new operator considered harmful.

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

## Usage

BuildFactory provide a factory of italian cars. Then italian builder build a particular model of car. This model must have 4 wheels.

```go
fca, _ := BuildFactory(ItalianType)
m, _ := fca.Build(FerrariModel)
car, err := m.(Vehicle)
if model.NumOfWheels() != 4 {
  panic("Ferrari shoud have 4 wheels")
}
```

BuildFactory provide a factory of italian cars. Then italian builder build a particular model of car. This model must have 5 wheels.

```go
fca, _ := BuildFactory(ItalianType)
model, _ := fca.Build(CarWithFiveWheelModel)
if model.NumOfWheels() != 5 {
  panic("the car should have 5 wheels")
}
```
